package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/alserov/device-shop/gateway/internal/broker"
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/config"
	"github.com/alserov/device-shop/gateway/internal/controller"
	"github.com/alserov/device-shop/gateway/internal/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
	log *slog.Logger

	server   *server
	services *controller.Services

	cacheAddr string
	broker    *broker.Broker
}

type App interface {
	MustStart()
}

type server struct {
	port   int
	router *gin.Engine
	server *http.Server
}

func New(cfg *config.Config) App {
	router := gin.New()
	return &app{
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
					Users:   cfg.Broker.Topics.Metrics.UsersAmount,
					Orders:  cfg.Broker.Topics.Metrics.Orders,
					Latency: cfg.Broker.Topics.Metrics.Latency,
				},
			},
		},
		log:       logger.MustSetupLogger(cfg.Env),
		cacheAddr: cfg.Cache.Addr,
		services: &controller.Services{
			OrderAddr:  cfg.Services.OrderAddr,
			CollAddr:   cfg.Services.CollAddr,
			UserAddr:   cfg.Services.UserAddr,
			DeviceAddr: cfg.Services.DeviceAddr,
		},
	}
}

const (
	kafkaClientID = "API_GATEWAY"
)

func (a *app) MustStart() {
	defer a.recover()
	a.log.Info("starting app", slog.Int("port", a.server.port))

	redisCl := cache.MustConnect(a.cacheAddr)
	a.log.Info("cache connected")
	cache := cache.NewRepo(redisCl)

	producer, err := broker.NewProducer([]string{a.broker.Addr}, kafkaClientID)
	if err != nil {
		panic("failed to init producer: " + err.Error())
	}

	metricsProducer := broker.NewMetricsProducer(producer, a.broker.Topics.Metrics)

	ctrl, closeConns := controller.NewController(&controller.Ctrl{
		Topics:          a.broker.Topics,
		Services:        a.services,
		Log:             a.log,
		Cache:           cache,
		MetricsProducer: metricsProducer,
	})
	defer closeConns()

	controller.LoadRoutes(a.server.router, ctrl)

	chStop := make(chan os.Signal)
	signal.Notify(chStop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		a.log.Info("app is running")
		if err = a.server.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("app has stopped due to the error: " + err.Error())
		}
	}()

	sign := <-chStop
	a.stop(sign)
}

func (a *app) stop(sign os.Signal) {
	if err := a.server.server.Shutdown(context.Background()); err != nil {
		panic("failed to shutdown server: " + err.Error())
	}
	a.log.Info("app was stopped due to: ", sign.String())
}

func (a *app) recover() {
	err := recover()
	if err != nil {
		a.log.Error("app was stopped due to panic", slog.Any("error", err))
	}
}
