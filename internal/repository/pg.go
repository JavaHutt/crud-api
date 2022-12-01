package repository

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPostgresDB() (*bun.DB, error) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("localhost:5432"),
		pgdriver.WithTLSConfig(nil),
		pgdriver.WithUser("postgres"),
		pgdriver.WithPassword("postgres"),
		pgdriver.WithDatabase("crud"),
		pgdriver.WithApplicationName("crud-api"),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	sqldb := sql.OpenDB(pgconn)
	pg := bun.NewDB(sqldb, pgdialect.New())
	if err := pg.Ping(); err != nil {
		return nil, err
	}
	return pg, nil
}
