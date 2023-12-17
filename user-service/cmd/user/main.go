package main

import (
	"github.com/alserov/device-shop/user-service/internal/app"
	"github.com/alserov/device-shop/user-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg)
	application.MustStart()
}
