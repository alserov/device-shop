package app

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"github.com/alserov/device-shop/order-service/internal/service"
	"github.com/alserov/device-shop/proto/gen"
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

	ordersDB, err := postgres.Connect(a.ordersDsn)
	if err != nil {
		return err
	}

	devicesDB, err := postgres.Connect(a.devicesDsn)
	if err != nil {
		return err
	}

	usersDB, err := postgres.Connect(a.usersDsn)
	if err != nil {
		return err
	}

	pb.RegisterOrdersServer(s, service.New(ordersDB, devicesDB, usersDB))

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
