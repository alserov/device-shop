package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type App struct {
	port           int
	host           string
	connType       string
	ordersDsn      string
	devicesDsn     string
	usersDsn       string
	collectionsUri string
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
		ordersDsn: fmt.Sprintf("host=%s port=%s user=%s password=%v dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSLMODE"),
		),
		devicesDsn: fmt.Sprintf("host=%s port=%s user=%s password=%v dbname=%s sslmode=%s",
			os.Getenv("DB_DEVICES_HOST"),
			os.Getenv("DB_DEVICES_PORT"),
			os.Getenv("DB_DEVICES_USER"),
			os.Getenv("DB_DEVICES_PASSWORD"),
			os.Getenv("DB_DEVICES_NAME"),
			os.Getenv("DB_DEVICES_SSLMODE"),
		),
		usersDsn: fmt.Sprintf("host=%s port=%s user=%s password=%v dbname=%s sslmode=%s",
			os.Getenv("DB_USERS_HOST"),
			os.Getenv("DB_USERS_PORT"),
			os.Getenv("DB_USERS_USER"),
			os.Getenv("DB_USERS_PASSWORD"),
			os.Getenv("DB_USERS_NAME"),
			os.Getenv("DB_USERS_SSLMODE"),
		),
		collectionsUri: os.Getenv("MONGO_URI"),
	}

	return a, nil
}
