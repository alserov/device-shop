package main

import (
	"github.com/alserov/device-shop/manager-service/internal/app"
	"github.com/alserov/device-shop/manager-service/internal/config"
	"github.com/alserov/device-shop/manager-service/internal/logger"
)

func main() {
	cfg := config.MustLoadEmail()

	log := logger.MustSetupLogger(cfg.Env)

	application := app.NewEmailApp(cfg, log)
	application.MustStartEmail()
}
