package postgres

import (
	"database/sql"
	"github.com/alserov/device-shop/user-service/internal/db/postgres/migrations"
	_ "github.com/lib/pq"
)

func MustConnect(dsn string) *sql.DB {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to open db: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping db: " + err.Error())
	}

	if err = migrations.Migrate(conn); err != nil {
		panic("failed to migrate schemas: " + err.Error())
	}

	return conn
}
