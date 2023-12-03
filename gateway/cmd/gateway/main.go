package main

import (
	"github.com/alserov/device-shop/gateway/internal/app"
	"github.com/alserov/device-shop/gateway/internal/config"
	"github.com/alserov/device-shop/gateway/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.MustSetupLogger(cfg.Env)

	application := app.New(cfg, log)
}
