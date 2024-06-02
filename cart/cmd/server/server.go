package main

import (
	"log/slog"
	"net"
	"net/http"
	"route256/cart/internal/app/server"
	"route256/cart/internal/pkg/cart/repository"
	"route256/cart/internal/pkg/cart/service/item/add_product"
	"route256/cart/internal/pkg/cart/service/item/delete_item"
	"route256/cart/internal/pkg/cart/service/list/clear_cart"
	"route256/cart/internal/pkg/cart/service/list/get_cart"
	"route256/cart/internal/pkg/middleware"
)

func main() {

	conn, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		slog.Error(err.Error())
	}
	defer conn.Close()

	cartRepository := repository.NewRepository()
	cartServer := server.NewServer()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", cartServer.AddItem(add_product.New(cartRepository)))
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", cartServer.DeleteItem(delete_item.New(cartRepository)))
	mux.HandleFunc("DELETE /user/{user_id}/cart", cartServer.DeleteItemsByUserID(clear_cart.New(cartRepository)))
	mux.HandleFunc("GET /user/{user_id}/cart/list", cartServer.GetItemsByUserID(get_cart.New(cartRepository)))

	handle := middleware.Logging(mux)
	handle = middleware.PanicRecovery(handle)

	err = http.Serve(conn, handle)
	if err != nil {
		slog.Error(err.Error())
	}
}
