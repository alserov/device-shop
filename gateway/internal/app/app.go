package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/config"
	"github.com/alserov/device-shop/gateway/internal/controller"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	port      int
	server    *http.Server
	log       *slog.Logger
	cacheAddr string
	router    *gin.Engine
	services  *config.Services
}

func New(cfg *config.Config, log *slog.Logger) *App {
	router := gin.New()
	return &App{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			WriteTimeout: cfg.Timeout * time.Second,
			ReadTimeout:  cfg.Timeout * time.Second,
			Handler:      router,
		},
		log:       log,
		router:    router,
		cacheAddr: cfg.Cache.Addr,
		services:  &cfg.Services,
		port:      cfg.Port,
	}
}

func (a *App) MustStart() {
	cl := cache.MustConnect(a.cacheAddr)

	controller.LoadRoutes(a.router, controller.NewController(cl, a.log, a.services))

	a.log.Info("server has started", slog.Int("port", a.port))
	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("failed to start server: " + err.Error())
	}
}

func (a *App) Stop() {
	if err := a.server.Shutdown(context.Background()); err != nil {
		panic("failed to shutdown server: " + err.Error())
	}
}
