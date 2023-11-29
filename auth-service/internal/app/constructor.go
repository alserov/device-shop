package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type App interface {
	Start(ctx context.Context) error
}

type app struct {
	port            int
	host            string
	connType        string
	postgresDsn     string
	emailTopic      string
	emailBrokerAddr string
}

const (
	DEFAULT_PORT = 8001
	DEFAULT_HOST = "localhost"
)

func New() (App, error) {
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

	a := &app{
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
		emailTopic:      os.Getenv("TOPIC"),
		emailBrokerAddr: os.Getenv("BROKER_ADDR"),
	}

	return a, nil
}
