package app

import (
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
	port    int
	timeout time.Duration
	router  *gin.Engine

	log *slog.Logger

	cacheAddr string
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		timeout: cfg.Timeout,
		//router:  gin.Default(),
		router: gin.New(),
		port:   cfg.Port,
		log:    log,
	}
}

func (a *App) MustStart() {
	cl := cache.MustConnect(a.cacheAddr)

	controller.LoadRoutes(a.router, controller.NewController(cl, a.log))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.port),
		Handler:      a.router,
		WriteTimeout: time.Duration(a.timeout) * time.Second,
		ReadTimeout:  time.Duration(a.timeout) * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
