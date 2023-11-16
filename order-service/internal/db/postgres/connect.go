package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	log.Println("postgres connected")
	return nil, nil
}
