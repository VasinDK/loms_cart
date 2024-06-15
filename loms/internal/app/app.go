package app

import (
	"fmt"
	"log/slog"
	"net"
	"route256/loms/internal/grpc_handlers"
	"route256/loms/internal/pkg/config"
	"route256/loms/internal/repositories/order"
	"route256/loms/internal/repositories/stock"
	"route256/loms/internal/service"
	"route256/loms/pkg/api/loms/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run - запускает grpc сервер
func Run(config *config.Config) {
	order := order.New()
	stock := stock.New()

	grpcService := service.New(
		order,
		stock,
	)

	grpcHandlers := grpc_handlers.New(grpcService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.GetPort()))
	if err != nil {
		panic(err.Error())
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	loms.RegisterLomsServer(grpcServer, grpcHandlers)

	slog.Info("server grpc listening at", lis.Addr())

	// go func() {
	if err := grpcServer.Serve(lis); err != nil {
		slog.Error("failed to grpc serve:", err)
	}
	// }()
}
