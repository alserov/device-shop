package controller

import (
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"os"
)

type Handler interface {
	Collectioner
	Devicer
	Adminer
	Auther
	Orderer
	Balancer
}

type handler struct {
	cache      cache.Repository
	logger     *logrus.Logger
	userAddr   string
	deviceAddr string
	orderAddr  string
}

func NewHandler(c *redis.Client, lg *logrus.Logger) Handler {
	var (
		userAddr   = os.Getenv("USER_ADDR")
		deviceAddr = os.Getenv("DEVICE_ADDR")
		orderAddr  = os.Getenv("ORDER_ADDR")
	)

	return &handler{
		cache:      cache.NewRepo(c),
		logger:     lg,
		userAddr:   userAddr,
		deviceAddr: deviceAddr,
		orderAddr:  orderAddr,
	}
}
