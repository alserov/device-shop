package app

import (
	"fmt"
	"github.com/alserov/admin-service/internal/config"
	"github.com/alserov/admin-service/internal/db/postgres"
	"github.com/alserov/admin-service/internal/server"
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
}

func New(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		port:       cfg.GRPC.Port,
		timeout:    cfg.GRPC.Timeout,
		log:        log,
		gRPCServer: grpc.NewServer(),
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
	a.log.Info("starting server on port", slog.Int("port", a.port))

	db := postgres.MustConnect(a.dbDsn)
	a.log.Info("db connected")

	server.Register(a.gRPCServer, db, a.log)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic("failed to start a server: " + err.Error())
	}

	if err = a.gRPCServer.Serve(l); err != nil {
		panic("app has stopped due to the error: " + err.Error())
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
