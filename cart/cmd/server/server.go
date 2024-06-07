package main

import (
	"route256/cart/internal/app"
	"route256/cart/internal/pkg/config"
)

func main() {
	config := config.New()
	app.Run(config)
}
