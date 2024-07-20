package main

import (
	"route256/notifier/internal/app"
	"route256/notifier/internal/pkg/config"
)

func main() {
	config := config.New()
	app.Run(config)
}
