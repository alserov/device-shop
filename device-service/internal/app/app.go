package app

import (
	"fmt"
	"github.com/alserov/device-shop/device-service/internal/broker"
	"github.com/alserov/device-shop/device-service/internal/broker/worker"
	"github.com/alserov/device-shop/device-service/internal/config"
	"github.com/alserov/device-shop/device-service/internal/db/postgres"
	"github.com/alserov/device-shop/device-service/internal/logger"
	"github.com/alserov/device-shop/device-service/internal/server"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type app struct {
	port  int
	dbDsn string

	log        *slog.Logger
	gRPCServer *grpc.Server
	broker     *broker.Broker
	services   *server.Services
}

type App interface {
	MustStart()
}

func New(cfg *config.Config) App {
	return &app{
		port: cfg.GRPC.Port,
		log:  logger.MustSetupLogger(cfg.Env),
		dbDsn: fmt.Sprintf("host=%s port=%d user=%s password=%v dbname=%s sslmode=%s",
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Name,
			cfg.DB.SSLMode,
		),
		services: &server.Services{
			CollectionAddr: cfg.Services.CollectionAddr,
		},
		gRPCServer: grpc.NewServer(),
		broker: &broker.Broker{
			Addr: cfg.Broker.Addr,
			Topics: &broker.Topics{
				Manager: broker.Topic{
					In:  cfg.Broker.WorkerTopicIn,
					Out: cfg.Broker.WorkerTopicOut,
				},
			},
		},
	}
}

func (a *app) MustStart() {
	defer a.recover()
	a.log.Info("starting app", slog.Int("port", a.port))

	db := postgres.MustConnect(a.dbDsn)
	a.log.Info("db connected")
	repo := postgres.NewRepo(db, a.log)

	w := worker.NewTxWorker(a.broker, repo, a.log)
	go w.MustStart()

	server.Register(&server.Server{
		GRPCServer: a.gRPCServer,
		Log:        a.log,
		Services:   a.services,
		Repo:       repo,
	})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to start a server: " + err.Error())
	}

	chStop := make(chan os.Signal)
	signal.Notify(chStop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		a.log.Info("app is running")
		if err = a.gRPCServer.Serve(l); err != nil {
			panic("app has stopped due to the error: " + err.Error())
		}
	}()

	sign := <-chStop
	a.stop(sign)
}

func (a *app) stop(sign os.Signal) {
	a.gRPCServer.GracefulStop()
	a.log.Info("app was stopped due to: ", sign.String())
}

func (a *app) recover() {
	err := recover()
	if err != nil {
		a.log.Error("app was stopped due to panic", slog.Any("error", err))
	}
}
