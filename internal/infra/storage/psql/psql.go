package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/illusory-server/accounts/internal/infra/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

//go:generate mockgen -package mockPsql -source psql.go -destination ../../../mock/psql/psql.go

type QueryExecutor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func ConnectPool(ctx context.Context, cfg *config.PostgreSQL) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		"disable",
	)
	conf, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, errors.Wrap(err, "pgxpool.ParseConfig")
	}

	conf.MinConns = int32(cfg.ConnectPoolMin) //nolint:gosec
	conf.MaxConns = int32(cfg.ConnectPoolMax) //nolint:gosec
	conf.MaxConnLifetime = time.Duration(cfg.ConnectTimeout) * time.Second
	conf.MaxConnIdleTime = time.Duration(cfg.ConnectIdle) * time.Second
	conf.HealthCheckPeriod = time.Duration(cfg.ConnectHealthCheck) * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, errors.Wrap(err, "pgxpool.NewWithConfig")
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "pool.Ping")
	}

	return pool, nil
}
