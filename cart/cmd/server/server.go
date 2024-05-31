package main

import (
	"log/slog"
	"net"
	"net/http"
	"route256/cart/internal/app/server"
	"route256/cart/internal/pkg/cart/repository"
	"route256/cart/internal/pkg/cart/service"
	"route256/cart/internal/pkg/middleware"
)

func main() {

	conn, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		slog.Error(err.Error())
	}
	defer conn.Close()

	cartRepository := repository.NewRepository()
	cartService := service.NewService(cartRepository)
	cartServer := server.NewServer(cartService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", cartServer.AddItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", cartServer.DeleteItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart", cartServer.DeleteItemsByUserID)
	mux.HandleFunc("GET /user/{user_id}/cart/list", cartServer.GetItemsByUserID)

	handle := middleware.Logging(mux)
	handle = middleware.PanicRecovery(handle)

	err = http.Serve(conn, handle)
	if err != nil {
		slog.Error(err.Error())
	}
}
