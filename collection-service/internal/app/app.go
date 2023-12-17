package app

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/collection-service/internal/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/alserov/device-shop/collection-service/internal/config"
	"github.com/alserov/device-shop/collection-service/internal/db/mongo"
	"github.com/alserov/device-shop/collection-service/internal/server"

	"google.golang.org/grpc"
	"log/slog"
	"net"
	"time"
)

type app struct {
	port  int
	dbDsn string

	gRPCServer *grpc.Server
	log        *slog.Logger
	services   *server.Services
}

type App interface {
	MustStart()
}

func New(cfg *config.Config) App {
	return &app{
		port:       cfg.GRPC.Port,
		gRPCServer: grpc.NewServer(),
		dbDsn:      cfg.DB.Uri,
		log:        logger.MustSetupLogger(cfg.Env),
		services: &server.Services{
			DeviceAddr: cfg.Services.DeviceAddr,
		},
	}
}

type Services struct {
	DeviceAddr string
}

func (a *app) MustStart() {
	defer a.recover()
	a.log.Info("starting app", slog.Int("port", a.port))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	db := mongo.MustConnect(ctx, a.dbDsn)
	a.log.Info("db connected")
	repo := mongo.NewRepo(db, a.log)

	server.Register(&server.Server{
		GRPCServer: a.gRPCServer,
		Log:        a.log,
		Services:   a.services,
		Repo:       repo,
	})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to start server: " + err.Error())
	}

	chStop := make(chan os.Signal)
	chErr := make(chan error)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		a.log.Info("app is running")
		if err = a.gRPCServer.Serve(l); err != nil {
			chErr <- err
		}
	}()

	sign := <-chStop
	a.stop(sign)
}

func (a *app) stop(signal os.Signal) {
	a.gRPCServer.GracefulStop()
	a.log.Info("app was stopped due to: " + signal.String())
}

func (a *app) recover() {
	err := recover()
	a.log.Error("app was stopped due to panic", slog.Any("error", err))
}
