package controller

import (
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/controller/handlers"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"os"
)

type Controller struct {
	cache              cache.Repository
	logger             *logrus.Logger
	authHandler        handlers.AuthHandler
	adminHandler       handlers.AdminHandler
	collectionsHandler handlers.CollectionsHandler
	devicesHandler     handlers.DevicesHandler
	orderHandler       handlers.OrdersHandler
	userHandler        handlers.UsersHandler
}

func NewController(c *redis.Client, lg *logrus.Logger) *Controller {
	var (
		userAddr   = os.Getenv("USER_ADDR")
		deviceAddr = os.Getenv("DEVICE_ADDR")
		orderAddr  = os.Getenv("ORDER_ADDR")
		authAddr   = os.Getenv("AUTH_ADDR")
	)

	return &Controller{
		cache:              cache.NewRepo(c),
		logger:             lg,
		adminHandler:       handlers.NewAdminHandler(deviceAddr, userAddr, lg),
		authHandler:        handlers.NewAuthHandler(authAddr, lg),
		collectionsHandler: handlers.NewCollectionsHandler(userAddr, lg),
		devicesHandler:     handlers.NewDevicesHandler(deviceAddr, cache.NewRepo(c), lg),
		orderHandler:       handlers.NewOrderHandler(orderAddr, lg),
		userHandler:        handlers.NewUserHandler(userAddr, lg),
	}
}
