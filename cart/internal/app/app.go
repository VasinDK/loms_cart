package app

import (
	"net/http"
	"route256/cart/internal/app/http_handlers"
	"route256/cart/internal/middleware"
	"route256/cart/internal/pkg/config"
	"route256/cart/internal/pkg/httpserver"
	"route256/cart/internal/repository"
	"route256/cart/internal/service/item/add_product"
	"route256/cart/internal/service/item/delete_item"
	"route256/cart/internal/service/list/clear_cart"
	"route256/cart/internal/service/list/get_cart"
)

func Run(config *config.Config) {
	cartRepository := repository.NewRepository(config)
	httpHandlers := http_handlers.New()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", httpHandlers.AddItem(add_product.New(cartRepository)))
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", httpHandlers.DeleteItem(delete_item.New(cartRepository)))
	mux.HandleFunc("DELETE /user/{user_id}/cart", httpHandlers.DeleteItemsByUserID(clear_cart.New(cartRepository)))
	mux.HandleFunc("GET /user/{user_id}/cart/list", httpHandlers.GetItemsByUserID(get_cart.New(cartRepository)))

	handle := middleware.Logging(mux)
	handle = middleware.PanicRecovery(handle)

	httpserver.Run(handle, config)
}
