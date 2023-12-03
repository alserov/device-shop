package app

import (
	"context"
	"fmt"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/alserov/device-shop/user-service/internal/db/postgres"
	"github.com/alserov/device-shop/user-service/internal/service"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
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

const (
	DEFAULT_PORT = 8001
	DEFAULT_HOST = "localhost"
)

func New() (*App, error) {
	var (
		port int
		err  error
	)

	if err = godotenv.Load(".env"); err != nil {
		return nil, err
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Println("SET DEFAULT VALUE FOR PORT: ", DEFAULT_PORT)
		port = DEFAULT_PORT
	} else {
		port, err = strconv.Atoi(portString)
		if err != nil {
			return nil, err
		}
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

	return a, err
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
