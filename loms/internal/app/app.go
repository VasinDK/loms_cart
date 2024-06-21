package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"route256/loms/internal/grpc_handlers"
	"route256/loms/internal/middleware"
	"route256/loms/internal/pkg/config"
	"route256/loms/internal/pkg/db"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/stock"
	"route256/loms/internal/service"
	"route256/loms/pkg/api/loms/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run - запускает grpc сервер
func Run(config *config.Config) {
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

	grpcHandlers := grpc_handlers.New(grpcService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GetPort()))
	if err != nil {
		panic(err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.Validate,
		),
	)

	reflection.Register(grpcServer)

	loms.RegisterLomsServer(grpcServer, grpcHandlers)

	slog.Info("server grpc listening at", lis.Addr())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("failed to grpc serve:", err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%v:%v", config.GetHost(), config.GetPort()),
		grpc.WithInsecure(),
	)
	if err != nil {
		slog.Error("Failed to dial server", err)
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()

	if err = loms.RegisterLomsHandler(context.Background(), gwmux, conn); err != nil {
		slog.Error("Failed to register gateway:", err)
	}

	slog.Info("Serving gRPC-Gateway on!")

	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.GetHttpPort()), gwmux); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
