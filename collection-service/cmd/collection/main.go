package main

import (
	"github.com/alserov/device-shop/collection-service/internal/app"
	"github.com/alserov/device-shop/collection-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg)
	application.MustStart()
}
