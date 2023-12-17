package app

import (
	"fmt"
	"github.com/alserov/device-shop/order-service/internal/broker"
	"github.com/alserov/device-shop/order-service/internal/config"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/logger"
	"github.com/alserov/device-shop/order-service/internal/server"
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
		log:        logger.MustSetupLogger(cfg.Env),
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
		broker: &broker.Broker{
			Addr: cfg.Broker.Addr,
			Topics: broker.Topics{
				Balance: broker.Topic{
					In:  cfg.Broker.UserTopicIn,
					Out: cfg.Broker.UserTopicOut,
				},
				Device: broker.Topic{
					In:  cfg.Broker.DeviceTopicIn,
					Out: cfg.Broker.DeviceTopicOut,
				},
				Collection: broker.Topic{
					In:  cfg.Broker.CollectionTopicIn,
					Out: cfg.Broker.CollectionTopicOut,
				},
			},
		},
		services: &server.Services{
			DeviceAddr: cfg.Services.DeviceAddr,
		},
	}
}

func (a *app) MustStart() {
	defer a.recover()
	a.log.Info("starting app", slog.Int("port", a.port))

	db := postgres.MustConnect(a.dbDsn)
	a.log.Info("db connected")
	repo := postgres.NewRepo(db, a.log)

	server.Register(&server.Server{
		Log:        a.log,
		Services:   a.services,
		GRPCServer: a.gRPCServer,
		Repo:       repo,
		Broker:     a.broker,
	})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	chStop := make(chan os.Signal)
	signal.Notify(chStop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		a.log.Info("app is running")
		if err = a.gRPCServer.Serve(l); err != nil {
			panic("failed to server: " + err.Error())
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
