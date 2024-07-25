package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"route256/loms/internal/grpc_handlers"
	"route256/loms/internal/middleware"
	"route256/loms/internal/model"
	"route256/loms/internal/pkg/config"
	"route256/loms/internal/pkg/db_shard"
	"route256/loms/internal/pkg/jaegertracing"
	"route256/loms/internal/pkg/logger"
	"route256/loms/internal/repositories/async_producer"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/stock"
	"syscall"

	"route256/loms/internal/service"
	"route256/loms/pkg/api/loms/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// Run - запускает grpc сервер
func Run(config *config.Config) {
	// Логер и контекст
	ctxStart := context.Background()
	logger.New()

	// Подключение к БД
	sm, err := db_shard.New(ctxStart, config)
	if err != nil {
		panic("dbShard.New " + err.Error())
	}
	defer sm.Close()

	// Репозитории
	order := order.New(sm)
	stock := stock.New(sm)

	// Брокер сообщений
	producer, err := async_producer.NewAsyncProducer(ctxStart, config)
	if err != nil {
		logger.Panicw(ctxStart, "err", err)
	}
	defer producer.AsyncProducer.AsyncClose()

	// Сервисы GRPC
	grpcService := service.New(
		order,
		stock,
		producer,
	)

	// Tracing
	tp, err := jaegertracing.New(config, model.ServiceName)
	if err != nil {
		logger.Panicw(ctxStart, "err", err)
	}

	// Обработчики GRPC
	grpcHandlers := grpc_handlers.New(grpcService)

	// Слушатель GRPC
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GetPort()))
	if err != nil {
		panic(err.Error())
	}

	// Сервер GRPC
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			middleware.Validate,
		),
	)

	// Список методов
	reflection.Register(grpcServer)

	// Регистрация gRPC сервера
	loms.RegisterLomsServer(grpcServer, grpcHandlers)

	logger.Infow(ctxStart, "server grpc listening at")

	// http gateway GRPC
	conn, err := grpc.NewClient(
		fmt.Sprintf("%v:%v", config.GetHost(), config.GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)

	if err != nil {
		logger.Errorw(ctxStart, "Failed to dial server", "err", err)
	}
	defer conn.Close()

	// Роутер gateway http
	gwmux := runtime.NewServeMux()

	// Регистрация http gateway GRPC
	if err = loms.RegisterLomsHandler(ctxStart, gwmux, conn); err != nil {
		logger.Errorw(ctxStart, "Failed to register gateway", "err", err)
	}

	logger.Infow(ctxStart, "Serving gRPC-Gateway on!")

	// Роутер http
	mux := http.NewServeMux()

	mux.Handle("/", gwmux)
	mux.Handle("GET /metrics", promhttp.Handler())

	go func() {
		// Запуск gRPC сервера
		if err := grpcServer.Serve(lis); err != nil {
			logger.Errorw(ctxStart, "failed to grpc serve", "err", err)
		}
	}()

	go func() {
		// Запуск http gateway gRPC сервера
		if err := http.ListenAndServe(fmt.Sprintf(":%v", config.GetHttpPort()), mux); err != nil {
			logger.Errorw(ctxStart, "Failed to start HTTP server", "err", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	logger.Infow(ctxStart, "A stop signal has been received")

	// Shutdown
	grpcServer.GracefulStop()

	if err = tp.Shutdown(ctxStart); err != nil {
		logger.Errorw(ctxStart, "Failed to shutdown TracerProvider", "err", err)
	}

	producer.AsyncProducer.AsyncClose()
}
