package server

import (
	"database/sql"
	"github.com/alserov/device-shop/auth-service/internal/service"
	"github.com/alserov/device-shop/proto/gen/auth"
	"google.golang.org/grpc"
	"log/slog"
)

func Register(s *grpc.Server, db *sql.DB, log *slog.Logger, kafkaTopic string, kafkaBrokerAddr string) {
	auth.RegisterAuthServer(s, &server{
		auth: service.NewService(db, kafkaTopic, kafkaBrokerAddr),
		log:  log,
	})
}

type server struct {
	auth.UnimplementedAuthServer
	auth service.Service
	log  *slog.Logger
}
