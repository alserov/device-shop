package postgres

import (
	"database/sql"
	"errors"
	"github.com/alserov/device-shop/device-service/internal/db/postgres/migrations"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
	"log"
)

func MustConnect(dsn string) *sql.DB {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to open db: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping db: " + err.Error())
	}

	if err = migrations.Migrate(conn); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic("failed to migrate: " + err.Error())
	}

	log.Println("postgres connected")
	return conn
}
