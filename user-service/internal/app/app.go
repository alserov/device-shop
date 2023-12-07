package app

import (
	"fmt"
	"github.com/alserov/device-shop/user-service/internal/broker"
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
	log        *slog.Logger
	timeout    time.Duration
	dbDsn      string
	gRPCServer *grpc.Server
	kafka      kafka
}

type kafka struct {
	brokerAddr     string
	emailTopic     string
	workerInTopic  string
	workerOutTopic string
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		port:       cfg.GRPC.Port,
		timeout:    cfg.GRPC.Timeout,
		log:        log,
		gRPCServer: grpc.NewServer(),
		kafka: kafka{
			emailTopic:     cfg.Kafka.EmailTopic,
			workerInTopic:  cfg.Kafka.WorkerInTopic,
			workerOutTopic: cfg.Kafka.WorkerOutTopic,
			brokerAddr:     cfg.Kafka.BrokerAddr,
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

	worker := broker.NewTxWorker(a.kafka.brokerAddr, a.kafka.workerInTopic, a.kafka.workerOutTopic, db, a.log)
	go worker.MustStart()

	server.Register(&server.Server{
		GRPCServer: a.gRPCServer,
		DB:         db,
		Log:        a.log,
		BrokerAddr: a.kafka.brokerAddr,
		EmailTopic: a.kafka.emailTopic,
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
