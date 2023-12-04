package controller

import (
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/config"
	"github.com/alserov/device-shop/gateway/internal/controller/handlers"
	"github.com/go-redis/redis"
	"log/slog"
)

type Controller struct {
	cache              cache.Repository
	logger             *slog.Logger
	authHandler        handlers.AuthHandler
	adminHandler       handlers.AdminHandler
	collectionsHandler handlers.CollectionsHandler
	devicesHandler     handlers.DevicesHandler
	orderHandler       handlers.OrdersHandler
	userHandler        handlers.UsersHandler
}

func NewController(c *redis.Client, lg *slog.Logger, services *config.Services) *Controller {
	return &Controller{
		cache:              cache.NewRepo(c),
		logger:             lg,
		adminHandler:       handlers.NewAdminHandler(services.Device.Addr, services.User.Addr, lg),
		authHandler:        handlers.NewAuthHandler(services.User.Addr, lg),
		collectionsHandler: handlers.NewCollectionsHandler(services.Coll.Addr, lg),
		devicesHandler:     handlers.NewDevicesHandler(services.Device.Addr, cache.NewRepo(c), lg),
		orderHandler:       handlers.NewOrderHandler(services.Order.Addr, lg),
		userHandler:        handlers.NewUserHandler(services.User.Addr, lg),
	}
}
