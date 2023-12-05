package app

import (
	"github.com/alserov/device-shop/order-service/internal/db/postgres"
	"google.golang.org/grpc"
	"log/slog"
	"time"
)

type App struct {
	port       int
	timeout    time.Duration
	dbDsn      string
	log        *slog.Logger
	gRPCServer *grpc.Server
	services   services
}

type services struct {
	CollectionAddr string
	DeviceAddr     string
	UserAddr       string
}

fmt.Sprintf("host=%s port=%s user=%s password=%v dbname=%s sslmode=%s",
os.Getenv("DB_HOST"),
os.Getenv("DB_PORT"),
os.Getenv("DB_USER"),
os.Getenv("DB_PASSWORD"),
os.Getenv("DB_NAME"),
os.Getenv("DB_SSLMODE"),
)

func New() *App {

}

func (a *App) MustStart()  {
	db := postgres.MustConnect(a.dbDsn)

	
}
