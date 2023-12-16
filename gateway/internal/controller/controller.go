package controller

import (
	"github.com/alserov/device-shop/gateway/internal/broker"
	"github.com/alserov/device-shop/gateway/internal/config"
	"github.com/alserov/device-shop/gateway/internal/controller/handlers"
	"github.com/alserov/device-shop/gateway/internal/services"
	"github.com/go-redis/redis"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"log/slog"
)

type Controller struct {
	metricsReg prometheus.Registerer

	authHandler       handlers.AuthHandler
	adminHandler      handlers.AdminHandler
	collectionHandler handlers.CollectionsHandler
	deviceHandler     handlers.DeviceHandler
	orderHandler      handlers.OrderHandler
	userHandler       handlers.UsersHandler
}

type Ctrl struct {
	Topics          *broker.Topics
	RedisClient     *redis.Client
	MetricsProducer broker.MetricsProducer
	Services        *config.Services
	Log             *slog.Logger
}

const servicesAmount = 4

func NewController(c *Ctrl) (*Controller, CloseConns) {
	conns := make([]*grpc.ClientConn, 0, servicesAmount)

	deviceClient, deviceConnection := services.NewDeviceClient(c.Services.Device.Addr)
	conns = append(conns, deviceConnection)

	orderClient, orderConnection := services.NewOrderClient(c.Services.Order.Addr)
	conns = append(conns, orderConnection)

	userClient, userConnection := services.NewUserClient(c.Services.User.Addr)
	conns = append(conns, userConnection)

	collectionClient, collectionConnection := services.NewCollectionClient(c.Services.Coll.Addr)
	conns = append(conns, collectionConnection)

	return &Controller{
			adminHandler:      handlers.NewAdminHandler(deviceClient, c.Log),
			authHandler:       handlers.NewAuthHandler(userClient, c.Log),
			collectionHandler: handlers.NewCollectionsHandler(collectionClient, c.Log),
			deviceHandler:     handlers.NewDeviceHandler(deviceClient, c.RedisClient, c.MetricsProducer, c.Log),
			orderHandler:      handlers.NewOrderHandler(orderClient, c.Log),
			userHandler:       handlers.NewUserHandler(userClient, c.Log),
		}, func() {
			CloseConnections(conns)
		}
}

type CloseConns func()

func CloseConnections(conns []*grpc.ClientConn) {
	for _, c := range conns {
		c.Close()
	}
}
