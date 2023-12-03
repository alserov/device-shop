package main

import (
	"github.com/alserov/device-shop/gateway/internal/app"
	"github.com/alserov/device-shop/gateway/internal/config"
	"github.com/alserov/device-shop/gateway/internal/logger"
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
	log.Info("server was stopped", slog.String("signal", sign.String()))
}
