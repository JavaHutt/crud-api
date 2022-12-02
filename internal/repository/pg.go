package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type config interface {
	AppName() string
	PostgresHost() string
	PostgresPort() string
	PostgresName() string
	PostgresUser() string
	PostgresPassword() string
	IdleTimeout() time.Duration
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
}

func NewPostgresDB(cfg config) (*bun.DB, error) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(fmt.Sprintf("%s:%s", cfg.PostgresHost(), cfg.PostgresPort())),
		pgdriver.WithTLSConfig(nil),
		pgdriver.WithApplicationName(cfg.AppName()),
		pgdriver.WithDatabase(cfg.PostgresName()),
		pgdriver.WithUser(cfg.PostgresUser()),
		pgdriver.WithPassword(cfg.PostgresPassword()),
		pgdriver.WithTimeout(cfg.IdleTimeout()),
		pgdriver.WithDialTimeout(cfg.IdleTimeout()),
		pgdriver.WithReadTimeout(cfg.ReadTimeout()),
		pgdriver.WithWriteTimeout(cfg.WriteTimeout()),
	)

	sqldb := sql.OpenDB(pgconn)
	pg := bun.NewDB(sqldb, pgdialect.New())
	if err := pg.Ping(); err != nil {
		return nil, err
	}
	return pg, nil
}
