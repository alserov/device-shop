package controller

import (
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	CollectionsHandler
	DeviceHandler
	AdminHandler
	AuthHandler
	OrderHandler
}

type handler struct {
	cache  cache.Repository
	logger *logrus.Logger
}

func NewHandler(c *redis.Client, lg *logrus.Logger) Handler {
	return &handler{
		cache:  cache.NewRepo(c),
		logger: lg,
	}
}
