package app

import (
	"fmt"
	"github.com/alserov/device-shop/user-service/internal/broker"
	"github.com/alserov/device-shop/user-service/internal/broker/worker"
	"github.com/alserov/device-shop/user-service/internal/config"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"github.com/alserov/device-shop/user-service/internal/server"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"time"
)

type App struct {
	port       int
	timeout    time.Duration
	log        *slog.Logger
	dbDsn      string
	gRPCServer *grpc.Server
	broker     *broker.Broker
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		port:       cfg.GRPC.Port,
		timeout:    cfg.GRPC.Timeout,
		log:        log,
		gRPCServer: grpc.NewServer(),
		broker: &broker.Broker{
			BrokerAddr: cfg.Kafka.BrokerAddr,
			Topics: broker.Topics{
				Email: cfg.Kafka.EmailTopic,
				Manager: broker.Topic{
					In:  cfg.Kafka.WorkerTopicIn,
					Out: cfg.Kafka.WorkerTopicOut,
				},
			},
		},
		dbDsn: fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=%s",
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Name,
			cfg.DB.SSLMode,
		),
	}
}

func (a *App) MustStart() {
	a.log.Info("starting app", slog.Int("port", a.port))

	db := postgres.MustConnect(a.dbDsn)
	a.log.Info("db connected")

	w := worker.NewWorker(a.broker, db, a.log)
	go w.MustStart()

	server.Register(&server.Server{
		GRPCServer: a.gRPCServer,
		DB:         db,
		Log:        a.log,
		BrokerAddr: a.broker.BrokerAddr,
		EmailTopic: a.broker.Topics.Email,
	})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to create listener: " + err.Error())
	}

	a.log.Info("app is running")
	if err = a.gRPCServer.Serve(l); err != nil {
		panic("failed to start the server: " + err.Error())
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
