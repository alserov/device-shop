package app

import (
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/controller"
	"github.com/alserov/device-shop/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

type App struct {
	host    string
	port    int
	timeout int
	router  *gin.Engine
	handler controller.Handler
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

	l, err := logger.New()
	if err != nil {
		log.Println("failed to initialize logger: ", err)
		return nil, err
	}

	a := &App{
		timeout: timeout,
		//router:  gin.Default(),
		router:  gin.New(),
		handler: controller.NewHandler(cl, l),
		port:    port,
		host:    host,
	}

	controller.LoadRoutes(a.router, a.handler)

	return a, nil
}
