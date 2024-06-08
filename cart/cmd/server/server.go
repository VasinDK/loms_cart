package main

import (
	"route256/cart/internal/app"
	"route256/cart/internal/pkg/config"
)

// main - точка входа в приложение
func main() {
	config := config.New()
	app.Run(config)
}
