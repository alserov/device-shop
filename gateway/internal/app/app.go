package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/alserov/device-shop/gateway/internal/broker"
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/config"
	"github.com/alserov/device-shop/gateway/internal/controller"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	log *slog.Logger

	server *server

	cacheAddr string
	broker    *broker.Broker
	services  *config.Services
}

type server struct {
	port   int
	router *gin.Engine
	server *http.Server
}

func New(cfg *config.Config, log *slog.Logger) *App {
	router := gin.New()
	return &App{
		server: &server{
			server: &http.Server{
				Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
				WriteTimeout: cfg.Server.Timeout * time.Second,
				ReadTimeout:  cfg.Server.Timeout * time.Second,
				Handler:      router,
			},
			router: router,
			port:   cfg.Server.Port,
		},
		broker: &broker.Broker{
			Addr: cfg.Broker.Addr,
			Topics: &broker.Topics{
				Metrics: &broker.Metrics{
					Request: &broker.RequestTopics{
						Total:      cfg.Broker.Topics.Metrics.Request.Total,
						Successful: cfg.Broker.Topics.Metrics.Request.Total,
					},
				},
			},
		},
		log:       log,
		cacheAddr: cfg.Cache.Addr,
		services:  &cfg.Services,
	}
}

const (
	kafkaClientID = "API_GATEWAY"
)

func (a *App) MustStart() {
	a.log.Info("starting app", slog.Int("port", a.server.port))

	cl := cache.MustConnect(a.cacheAddr)
	a.log.Info("cache connected")

	producer, err := broker.NewProducer([]string{a.broker.Addr}, kafkaClientID)
	if err != nil {
		panic("failed to init producer: " + err.Error())
	}

	controller.LoadRoutes(a.server.router, controller.NewController(&controller.C{
		Topics:      a.broker.Topics,
		RedisClient: cl,
		Producer:    producer,
		Services:    a.services,
		Log:         a.log,
	}))

	a.log.Info("app is running")
	if err = a.server.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("failed to start server: " + err.Error())
	}
}

func (a *App) Stop() {
	if err := a.server.server.Shutdown(context.Background()); err != nil {
		panic("failed to shutdown server: " + err.Error())
	}
}
