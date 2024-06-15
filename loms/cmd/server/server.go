package main

import (
	"route256/loms/internal/app"
	"route256/loms/internal/pkg/config"
)

// main - точка входа в приложение
func main() {
	config := config.New()
	app.Run(config)
}
