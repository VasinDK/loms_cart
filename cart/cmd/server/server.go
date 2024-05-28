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

	// добавить комменты и логирование

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", cartServer.AddItemCart)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", cartServer.DelItemCart)
	mux.HandleFunc("DELETE /user/{user_id}/cart", cartServer.CleaningCart)
	mux.HandleFunc("GET /user/{user_id}/cart", cartServer.GetCart)

	http.Serve(conn, mux)
}
