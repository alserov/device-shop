package app

import (
	"context"
	"fmt"
	"github.com/alserov/shop/gateway/internal/cache"
	"github.com/alserov/shop/gateway/internal/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type App struct {
	host    string
	port    int
	timeout int
	router  *gin.Engine
	handler controller.Handler
}

func (a *App) Start(ctx context.Context) error {
	log.Println("Starting API Gateway")
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.port),
		Handler:      a.router,
		WriteTimeout: time.Duration(a.timeout) * time.Second,
		ReadTimeout:  time.Duration(a.timeout) * time.Second,
	}

	chErr := make(chan error, 1)
	go func() {
		log.Println("API Gateway is working")
		if err := srv.ListenAndServe(); err != nil {
			chErr <- err
		}
	}()

	select {
	case <-ctx.Done():
		err := srv.Shutdown(ctx)
		if err != nil {
			return err
		}
		return nil
	case err := <-chErr:
		return err
	}
}

type Config struct {
	Server struct {
		Port    int    `yaml:"port"`
		Host    string `yaml:"host"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"server"`
}

const (
	DEFAULT_PORT    = 8001
	DEFAULT_HOST    = "localhost"
	DEFAULT_TIMEOUT = 10
)

func New() (*App, error) {
	host := os.Getenv("HOST")
	if host == "" {
		log.Println("SET DEFAULT VALUE FOR HOST: ", DEFAULT_HOST)
		host = DEFAULT_HOST
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	if port == 0 {
		log.Println("SET DEFAULT VALUE FOR PORT: ", DEFAULT_PORT)
		port = DEFAULT_PORT
	}

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		return nil, err
	}
	if timeout == 0 {
		log.Println("SET DEFAULT VALUE FOR TIMEOUT: ", DEFAULT_TIMEOUT)
		port = DEFAULT_PORT
	}

	cl, err := cache.Connect()
	if err != nil {
		return nil, err
	}

	a := &App{
		timeout: timeout,
		//router:  gin.Default(),
		router:  gin.New(),
		handler: controller.NewHandler(cl),
		port:    port,
		host:    host,
	}

	controller.LoadRoutes(a.router, a.handler)

	return a, nil
}
