package app

import (
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/controller"
	"github.com/alserov/device-shop/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type App struct {
	host    string
	port    int
	timeout int
	router  *gin.Engine
	handler *controller.Controller
}

const (
	DEFAULT_PORT    = 8001
	DEFAULT_HOST    = "localhost"
	DEFAULT_TIMEOUT = 10
)

func New() (*App, error) {
	var (
		port    int
		timeout int
		err     error
	)

	if err = godotenv.Load(".env"); err != nil {
		return nil, err
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Println("SET DEFAULT VALUE FOR PORT: ", DEFAULT_PORT)
		port = DEFAULT_PORT
	} else {
		port, err = strconv.Atoi(portString)
		if err != nil {
			return nil, err
		}
	}

	host := os.Getenv("HOST")
	if host == "" {
		log.Println("SET DEFAULT VALUE FOR HOST: ", DEFAULT_HOST)
		host = DEFAULT_HOST
	}

	timeoutString := os.Getenv("TIMEOUT")
	if timeoutString == "" {
		log.Println("SET DEFAULT VALUE FOR TIMEOUT: ", DEFAULT_TIMEOUT)
		timeout = DEFAULT_PORT
	} else {
		timeout, err = strconv.Atoi(timeoutString)
		if err != nil {
			return nil, err
		}
	}

	cl, err := cache.Connect()
	if err != nil {
		return nil, err
	}

	l, err := logger.New("internal")
	if err != nil {
		log.Println("failed to initialize logger: ", err)
		return nil, err
	}

	a := &App{
		timeout: timeout,
		//router:  gin.Default(),
		router:  gin.New(),
		handler: controller.NewController(cl, l),
		port:    port,
		host:    host,
	}

	controller.LoadRoutes(a.router, a.handler)

	return a, nil
}
