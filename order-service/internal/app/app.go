package app

import (
	"fmt"
	"github.com/alserov/device-shop/order-service/internal/broker"
	"github.com/alserov/device-shop/order-service/internal/config"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/server"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"time"
)

type App struct {
	port       int
	timeout    time.Duration
	dbDsn      string
	log        *slog.Logger
	gRPCServer *grpc.Server
	kafka      *broker.Broker
	services   services
}

type services struct {
	deviceAddr string
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		log:        log,
		port:       cfg.GRPC.Port,
		gRPCServer: grpc.NewServer(),
		dbDsn: fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=%s",
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Name,
			cfg.DB.Sslmode,
		),
		kafka: &broker.Broker{
			BrokerAddr: cfg.Kafka.BrokerAddr,
			Topics: broker.Topics{
				User: broker.Topic{
					In:  cfg.Kafka.UserTopicIn,
					Out: cfg.Kafka.UserTopicOut,
				},
				Device: broker.Topic{
					In:  cfg.Kafka.DeviceTopicIn,
					Out: cfg.Kafka.DeviceTopicOut,
				},
				Collection: broker.Topic{
					In:  cfg.Kafka.CollectionTopicIn,
					Out: cfg.Kafka.CollectionTopicOut,
				},
			},
		},
		services: services{
			deviceAddr: cfg.Services.DeviceAddr,
		},
	}
}

func (a *App) MustStart() {
	a.log.Info("starting app", slog.Int("port", a.port))

	db := postgres.MustConnect(a.dbDsn)
	a.log.Info("db connected")

	server.Register(&server.Server{
		Log:        a.log,
		DeviceAddr: a.services.deviceAddr,
		GRPCServer: a.gRPCServer,
		DB:         db,
		Broker:     a.kafka,
	})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	a.log.Info("app is running")
	if err = a.gRPCServer.Serve(l); err != nil {
		panic("failed to serve: " + err.Error())
	}
}

func (a *App) Stop() {
	a.gRPCServer.Stop()
}
