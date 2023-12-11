package main

import (
	"github.com/alserov/device-shop/metrics/internal/app"
	"github.com/alserov/device-shop/metrics/internal/config"
	"github.com/alserov/device-shop/metrics/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.MustSetupLogger(cfg.Env)

	application := app.New(cfg, log)
	application.MustStart()
}
