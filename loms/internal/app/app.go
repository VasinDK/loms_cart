package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"route256/loms/internal/grpc_handlers"
	"route256/loms/internal/middleware"
	"route256/loms/internal/model"
	"route256/loms/internal/pkg/config"
	"route256/loms/internal/pkg/db"
	"route256/loms/internal/pkg/jaegertracing"
	"route256/loms/internal/pkg/logger"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/stock"
	"route256/loms/internal/service"
	"route256/loms/pkg/api/loms/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run - запускает grpc сервер
func Run(config *config.Config) {
	ctxStart := context.Background()
	logger.New()

	connDB, err := db.NewConn(config)
	if err != nil {
		panic("db.NewConn " + err.Error())
	}
	defer connDB.Close()

	order := order.New(connDB)
	stock := stock.New(connDB)

	grpcService := service.New(
		order,
		stock,
	)

	tp, err := jaegertracing.New(config, model.ServiceName)
	if err != nil {
		logger.Panicw(ctxStart, "err", err)
	}

	grpcHandlers := grpc_handlers.New(grpcService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GetPort()))
	if err != nil {
		panic(err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
			middleware.Validate,
		),
	)

	reflection.Register(grpcServer)

	loms.RegisterLomsServer(grpcServer, grpcHandlers)

	logger.Infow(ctxStart, "server grpc listening at", lis.Addr())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Errorw(ctxStart, "failed to grpc serve", "err", err)
		}

		if err = tp.Shutdown(ctxStart); err != nil {
			logger.Errorw(ctxStart, "Failed to shutdown TracerProvider", "err", err)
		}
	}()

	conn, err := grpc.DialContext(
		ctxStart,
		fmt.Sprintf("%v:%v", config.GetHost(), config.GetPort()),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Errorw(ctxStart, "Failed to dial server", "err", err)
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()

	if err = loms.RegisterLomsHandler(ctxStart, gwmux, conn); err != nil {
		logger.Errorw(ctxStart, "Failed to register gateway", "err", err)
	}

	logger.Infow(ctxStart, "Serving gRPC-Gateway on!")

	mux := http.NewServeMux()

	mux.Handle("/", gwmux)
	mux.Handle("GET /metrics", promhttp.Handler())

	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.GetHttpPort()), mux); err != nil {
		logger.Errorw(ctxStart, "Failed to start HTTP server", "err", err)
	}
}
