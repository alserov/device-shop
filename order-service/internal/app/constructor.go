package app

import (
	"log"
	"os"
	"strconv"
)

type App struct {
	port     int
	host     string
	connType string
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
	}

	return a, nil
}
