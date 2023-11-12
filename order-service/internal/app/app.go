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
	"os"
	"strconv"
)

type App struct {
	port     int
	host     string
	connType string
}

func (a *App) Start(ctx context.Context) error {
	log.Println("starting service")
	lis, err := net.Listen(a.connType, fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	pg, err := postgres.Connect()
	if err != nil {
		return err
	}

	pb.RegisterOrdersServer(s, service.New(pg))

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

const (
	DEFAULT_PORT = 8001
	DEFAULT_HOST = "localhost"
)

func New() (*App, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	if port == 0 {
		log.Println("SET DEFAULT VALUE FOR PORT: ", DEFAULT_PORT)
		port = DEFAULT_PORT
	}

	host := os.Getenv("HOST")
	if host == "" {
		log.Println("SET DEFAULT VALUE FOR HOST: ", DEFAULT_HOST)
		host = DEFAULT_HOST
	}

	a := &App{
		port:     port,
		host:     host,
		connType: "tcp",
	}

	return a, nil
}
