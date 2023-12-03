package app

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/collection-service/internal/config"
	"github.com/alserov/device-shop/collection-service/internal/db/mongo"
	"github.com/alserov/device-shop/collection-service/internal/server"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"time"
)

type App struct {
	port       int
	gRPCServer *grpc.Server
	log        *slog.Logger
	timeout    time.Duration
	dbUri      string
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		port:       cfg.GRPC.Port,
		timeout:    cfg.GRPC.Timeout,
		gRPCServer: grpc.NewServer(),
		dbUri:      cfg.DB.Uri,
		log:        log,
	}
}

func (a *App) MustStart() {
	a.log.Info("starting server", slog.Int("port", a.port))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	db := mongo.MustConnect(ctx, a.dbUri)

	server.Register(a.gRPCServer, db, a.log)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to start server: " + err.Error())
	}

	if err = a.gRPCServer.Serve(l); err != nil {
		panic("app has stopped due to the error: " + err.Error())
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
