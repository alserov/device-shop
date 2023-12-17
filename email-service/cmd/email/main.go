package main

import (
	"github.com/alserov/device-shop/email-service/internal/app"
	"github.com/alserov/device-shop/email-service/internal/config"
)

func main() {
	cfg := config.MustLoadEmail()

	application := app.NewApp(cfg)
	application.MustStart()
}
