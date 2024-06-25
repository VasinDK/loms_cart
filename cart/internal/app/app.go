package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"route256/cart/internal/app/http_handlers"
	"route256/cart/internal/middleware"
	"route256/cart/internal/pkg/config"
	"route256/cart/internal/pkg/httpserver"
	"route256/cart/internal/repository"
	"route256/cart/internal/service/item/add_product"
	"route256/cart/internal/service/item/delete_item"
	"route256/cart/internal/service/list/checkout"
	"route256/cart/internal/service/list/clear_cart"
	"route256/cart/internal/service/list/get_cart"
	"syscall"

	"route256/cart/pkg/api/loms/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run - запускает сервер
func Run(config *config.Config) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", config.GetAddressStoreLoms(), config.GetPort()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error("Loms start", err)
		os.Exit(1)
	}

	clientLoms := loms.NewLomsClient(conn)

	cartRepository := repository.NewRepository(config, clientLoms)
	httpHandlers := http_handlers.New()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", httpHandlers.AddItem(add_product.New(cartRepository)))
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", httpHandlers.DeleteItem(delete_item.New(cartRepository)))
	mux.HandleFunc("DELETE /user/{user_id}/cart", httpHandlers.DeleteItemsByUserID(clear_cart.New(cartRepository)))
	mux.HandleFunc("GET /user/{user_id}/cart/list", httpHandlers.GetItemsByUserID(get_cart.New(cartRepository)))
	mux.HandleFunc("POST /user/cart/checkout", httpHandlers.Checkout(checkout.New(cartRepository)))

	handle := middleware.Logging(mux)
	handle = middleware.PanicRecovery(handle)

	server := httpserver.New(handle, config)
	go server.Run()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	server.GraceShutdown(ctx)

	slog.Info("the server is beautifully stopped")
}
