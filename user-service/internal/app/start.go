package app

import (
	"context"
	"fmt"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/alserov/device-shop/user-service/internal/db/mongo"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"github.com/alserov/device-shop/user-service/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func (a *App) Start(ctx context.Context) error {
	log.Println("starting service")
	lis, err := net.Listen(a.connType, fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	pg, err := postgres.Connect(a.postgresDsn)
	if err != nil {
		return err
	}
	mg, err := mongo.Connect(ctx, a.mongoUri)
	if err != nil {
		return err
	}

	pb.RegisterUsersServer(s, service.New(pg, mg))

	chErr := make(chan error, 1)
	go func() {
		log.Printf("Service is running on: %s:%d", a.host, a.port)
		if err = s.Serve(lis); err != nil {
			chErr <- err
		}
	}()

	select {
	case <-ctx.Done():
		s.GracefulStop()
		return nil
	case err = <-chErr:
		return err
	}
}
