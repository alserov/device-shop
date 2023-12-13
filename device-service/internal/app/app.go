package app

import (
	"fmt"
	"github.com/alserov/device-shop/device-service/internal/broker"
	"github.com/alserov/device-shop/device-service/internal/broker/worker"
	"github.com/alserov/device-shop/device-service/internal/config"
	"github.com/alserov/device-shop/device-service/internal/db/postgres"
	"github.com/alserov/device-shop/device-service/internal/server"
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
	broker     *broker.Broker
	services   Services
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		port:    cfg.GRPC.Port,
		timeout: cfg.GRPC.Timeout,
		log:     log,
		dbDsn: fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=%s",
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Name,
			cfg.DB.SSLMode,
		),
		services: Services{
			CollectionAddr: cfg.Services.CollectionAddr,
		},
		gRPCServer: grpc.NewServer(),
		broker: &broker.Broker{
			Addr: cfg.Broker.Addr,
			Topics: broker.Topics{
				Manager: broker.Topic{
					In:  cfg.Broker.WorkerTopicIn,
					Out: cfg.Broker.WorkerTopicOut,
				},
			},
		},
	}
}

type Services struct {
	CollectionAddr string
}

func (a *App) MustStart() {
	a.log.Info("starting app", slog.Int("port", a.port))

	db := postgres.MustConnect(a.dbDsn)
	a.log.Info("db connected")

	w := worker.NewTxWorker(a.broker, db, a.log)
	go w.MustStart()

	server.Register(a.gRPCServer, db, a.log, a.services.CollectionAddr)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to start a server: " + err.Error())
	}

	a.log.Info("app is running")
	if err = a.gRPCServer.Serve(l); err != nil {
		panic("app has stopped due to the error: " + err.Error())
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
