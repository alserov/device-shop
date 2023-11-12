package app

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/device-service/internal/db/mongo"
	"github.com/alserov/device-shop/device-service/internal/db/postgres"
	"github.com/alserov/device-shop/device-service/internal/service"
	pb "github.com/alserov/device-shop/proto/gen"
	"log"
	"net"
	"os"
	"strconv"
)

type App struct {
	port        int
	host        string
	connType    string
	postgresDsn string
	mongoUri    string
}

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

	pb.RegisterDevicesServer(s, service.New(pg, mg))

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
		postgresDsn: fmt.Sprintf("host=%s port=%s user=%s password=%v dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSLMODE"),
		),
		mongoUri: os.Getenv("MONGO_URI"),
	}

	return a, nil
}
