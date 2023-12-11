package app

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/metrics/internal/broker"
	"github.com/alserov/device-shop/metrics/internal/config"
	"github.com/alserov/device-shop/metrics/internal/metric"
	"github.com/alserov/device-shop/metrics/internal/workers"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
	log    *slog.Logger
	broker *broker.Broker

	port         int
	readTimeout  time.Duration
	writeTimeout time.Duration
	server       *http.Server
}

type App interface {
	MustStart()
}

func New(cfg *config.Config, log *slog.Logger) App {
	return &app{
		log: log,
		broker: &broker.Broker{
			Addr: cfg.Broker.Addr,
			Topics: &broker.Topics{
				Request: cfg.Broker.Topics.Request,
			},
		},
		port: cfg.Server.Port,
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", cfg.Server.Port),
		},
	}
}

func (a *app) MustStart() {
	a.log.Info("starting app", slog.Int("port", a.port))

	reg := prometheus.NewRegistry()

	counterWorker, counterMetric := workers.NewRequestWorker(a.broker.Addr, a.broker.Topics.Request)
	go func() {
		counterWorker.Start()
	}()

	metric.Setup(reg, counterMetric)

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			panic("failed to start server: " + err.Error())
		}
	}()

	chStop := make(chan os.Signal)
	signal.Notify(chStop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-chStop
	a.stop()
	a.log.Info("app was stopped", slog.String("signal", sign.String()))
}

func (a *app) stop() {
	if err := a.server.Shutdown(context.Background()); err != nil {
		panic("failed to shutdown server: " + err.Error())
	}
}
