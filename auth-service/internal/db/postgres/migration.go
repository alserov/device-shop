package postgres

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(dsn string) error {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	driver, err := pg.WithInstance(conn, &pg.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/postgres/migrations/",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
