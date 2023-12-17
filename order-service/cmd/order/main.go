package main

import (
	"github.com/alserov/device-shop/order-service/internal/app"
	"github.com/alserov/device-shop/order-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg)
	application.MustStart()
}
