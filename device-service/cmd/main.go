package main

import (
	"github.com/alserov/device-shop/device-service/internal/app"
	"github.com/alserov/device-shop/device-service/internal/config"
	"github.com/alserov/device-shop/device-service/internal/logger"
	"log/slog"
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
	log.Info("app was stopped", slog.String("signal", sign.String()))

}
