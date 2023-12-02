package postgres

import "database/sql"

func MustConnect(dsn string) *sql.DB {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping db: " + err.Error())
	}

	return conn
}
