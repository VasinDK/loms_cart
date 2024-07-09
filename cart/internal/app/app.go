package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"route256/cart/internal/app/http_handlers"
	"route256/cart/internal/middleware"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/config"
	"route256/cart/internal/pkg/httpserver"
	"route256/cart/internal/pkg/jaegertracing"
	"route256/cart/internal/pkg/logger"
	"route256/cart/internal/repository"
	"route256/cart/internal/service/item/add_product"
	"route256/cart/internal/service/item/delete_item"
	"route256/cart/internal/service/list/checkout"
	"route256/cart/internal/service/list/clear_cart"
	"route256/cart/internal/service/list/get_cart"
	"syscall"

	"route256/cart/pkg/api/loms/v1"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run - запускает сервер
func Run(config *config.Config) {
	ctxStart := context.Background()
	logger.New()

	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(
		fmt.Sprintf("%v:%v", config.GetAddressStoreLoms(), config.GetPortLoms()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)

	if err != nil {
		logger.Errorw(ctxStart, "grpc.NewClient", "err", err)
		os.Exit(1)
	}

	tp, err := jaegertracing.New(config, model.ServiceName)
	if err != nil {
		logger.Panicw(ctxStart, "err", err)
	}

	clientLoms := loms.NewLomsClient(conn)

	cartRepository := repository.NewRepository(config, clientLoms)
	httpHandlers := http_handlers.New()

	mux := http.NewServeMux()

	mux.Handle("GET /metrics", promhttp.Handler())
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", httpHandlers.AddItem(add_product.New(cartRepository)))
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", httpHandlers.DeleteItem(delete_item.New(cartRepository)))
	mux.HandleFunc("DELETE /user/{user_id}/cart", httpHandlers.DeleteItemsByUserID(clear_cart.New(cartRepository)))
	mux.HandleFunc("GET /user/{user_id}/cart/list", httpHandlers.GetItemsByUserID(get_cart.New(cartRepository)))
	mux.HandleFunc("POST /user/cart/checkout", httpHandlers.Checkout(checkout.New(cartRepository)))

	handle := middleware.Tracing(mux)
	handle = middleware.Logging(handle)
	handle = middleware.PanicRecovery(handle)

	server := httpserver.New(handle, config)
	errRun := server.Run()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-errRun:
		logger.Errorw(ctxStart, "server.Run", "err", err)
	case <-ctx.Done():
		logger.Infow(ctxStart, "signal.NotifyContext stop")
	}

	go tp.Shutdown(context.Background())
	server.GraceShutdown()

	logger.Infow(ctxStart, "the server is beautifully stopped")
}
