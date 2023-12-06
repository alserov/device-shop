package app

import (
	"fmt"
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
	kafka      kafka
}

type kafka struct {
	brokerAddr string
	topic      string
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		log:        log,
		port:       cfg.GRPC.Port,
		gRPCServer: grpc.NewServer(),
		dbDsn: fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=%s",
			cfg.Db.Host,
			cfg.Db.Port,
			cfg.Db.User,
			cfg.Db.Password,
			cfg.Db.Name,
			cfg.Db.Sslmode,
		),
		kafka: kafka{
			topic:      cfg.Kafka.Topic,
			brokerAddr: cfg.Kafka.BrokerAddr,
		},
	}
}

func (a *App) MustStart() {
	a.log.Info("starting app", slog.Int("port", a.port))

	db := postgres.MustConnect(a.dbDsn)
	a.log.Info("db connected")

	server.Register(&server.Server{
		Log:        a.log,
		GRPCServer: a.gRPCServer,
		DB:         db,
		Kafka: &server.Kafka{
			BrokerAddr: a.kafka.brokerAddr,
			OrderTopic: a.kafka.topic,
		},
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
