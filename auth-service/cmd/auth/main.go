package main

import (
	"github.com/alserov/device-shop/auth-service/internal/app"
	"github.com/alserov/device-shop/auth-service/internal/config"
	"github.com/alserov/device-shop/auth-service/internal/logger"
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
	signal.Notify(chStop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-chStop
	application.Stop()
	log.Info("application was stopped due to: " + sign.String())
}
