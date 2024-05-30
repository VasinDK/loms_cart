package main

import (
	"log"
	"net"
	"net/http"
	"route256/cart/internal/app/server"
	"route256/cart/internal/pkg/cart/repository"
	"route256/cart/internal/pkg/cart/service"
)

func main() {

	conn, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Panic(err)
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

	http.Serve(conn, mux)
}
