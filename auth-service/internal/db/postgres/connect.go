package postgres

import (
	"database/sql"
	"github.com/alserov/device-shop/auth-service/internal/db/postgres/migrations"
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

	if err = migrations.Migrate(conn); err != nil {
		return nil, err
	}

	return conn, nil
}
