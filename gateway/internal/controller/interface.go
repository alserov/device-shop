package controller

import (
	"github.com/alserov/shop/gateway/internal/cache"
	"github.com/go-redis/redis"
)

type Handler interface {
	CollectionsHandler
	DeviceHandler
	AdminHandler
	AuthHandler
	OrderHandler
}

type handler struct {
	cache cache.Repository
}

func NewHandler(c *redis.Client) Handler {
	return &handler{
		cache: cache.NewRepo(c),
	}
}
