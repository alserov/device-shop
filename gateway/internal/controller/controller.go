package controller

import (
	"github.com/IBM/sarama"
	"github.com/alserov/device-shop/gateway/internal/broker"
	"github.com/alserov/device-shop/gateway/internal/config"
	"github.com/alserov/device-shop/gateway/internal/controller/handlers"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/go-redis/redis"
	"log/slog"
)

type Controller struct {
	metricsReg prometheus.Registerer

	authHandler       handlers.AuthHandler
	adminHandler      handlers.AdminHandler
	collectionHandler handlers.CollectionsHandler
	deviceHandler     handlers.DeviceHandler
	orderHandler      handlers.OrdersHandler
	userHandler       handlers.UsersHandler
}

type C struct {
	Topics      *broker.Topics
	RedisClient *redis.Client
	Producer    sarama.SyncProducer
	Services    *config.Services
	Log         *slog.Logger
}

func NewController(c *C) *Controller {

	deviceHandler := &handlers.DeviceH{
		DeviceAddr:   c.Services.Device.Addr,
		Producer:     c.Producer,
		Log:          c.Log,
		RequestTopic: c.Topics.Metrics.Request,
		RedisClient:  c.RedisClient,
	}

	adminHandler := &handlers.AdminH{
		DeviceAddr: c.Services.Device.Addr,
		Log:        c.Log,
	}

	authHandler := &handlers.AuthH{
		AuthAddr: c.Services.User.Addr,
		Log:      c.Log,
	}

	collectionHandler := &handlers.CollectionH{
		UserAddr: c.Services.Coll.Addr,
		Log:      c.Log,
	}

	orderHandler := &handlers.OrderH{
		OrderAddr: c.Services.Order.Addr,
		Log:       c.Log,
	}

	userHandler := &handlers.UserH{
		UserAddr: c.Services.User.Addr,
		Log:      c.Log,
	}

	return &Controller{
		adminHandler:      handlers.NewAdminHandler(adminHandler),
		authHandler:       handlers.NewAuthHandler(authHandler),
		collectionHandler: handlers.NewCollectionsHandler(collectionHandler),
		deviceHandler:     handlers.NewDeviceHandler(deviceHandler),
		orderHandler:      handlers.NewOrderHandler(orderHandler),
		userHandler:       handlers.NewUserHandler(userHandler),
	}
}
