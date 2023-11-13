package app

import (
	"fmt"
	"log"
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
