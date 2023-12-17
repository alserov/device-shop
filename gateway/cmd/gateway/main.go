package main

import (
	"github.com/alserov/device-shop/gateway/internal/app"
	"github.com/alserov/device-shop/gateway/internal/config"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg)
	application.MustStart()
}
