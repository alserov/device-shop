package main

import (
	"github.com/alserov/device-shop/collection-service/internal/app"
	"github.com/alserov/device-shop/collection-service/internal/config"
	"github.com/alserov/device-shop/collection-service/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.MustSetupLogger(cfg.Env)

	application := app.New(cfg, log)
	go application.MustStart()

	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-chStop
	application.Stop()
	log.Info("app was stopped due to: " + sign.String())
}
