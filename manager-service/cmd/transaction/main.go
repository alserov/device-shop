package main

import (
	"github.com/alserov/device-shop/manager-service/internal/app"
	"github.com/alserov/device-shop/manager-service/internal/config"
	"github.com/alserov/device-shop/manager-service/internal/logger"
)

func main() {
	cfg := config.MustLoadTransaction()

	log := logger.MustSetupLogger(cfg.Env)

	application := app.NewTxApp(cfg, log)
	application.MustStart()
}
