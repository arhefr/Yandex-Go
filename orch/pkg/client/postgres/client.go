package postgres

import (
	"context"
	"fmt"
	"time"

	log "github.com/arhefr/Yandex-Go/orch/pkg/logger"
	repeatible "github.com/arhefr/Yandex-Go/orch/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(logger *log.Logger, ctx context.Context, cfg DBConfig, build string) (pool *pgxpool.Pool, err error) {
	time.Sleep(3 * time.Second)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	err = repeatible.DoWithTries(func() error {
		logger.Info("Trying connect to postgress...")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, connStr)
		if err != nil {
			return fmt.Errorf("postgres: NewClient: error failed to connect to postgres")
		}

		return nil
	}, cfg.MaxAtmps, time.Duration(cfg.DelayAtmpsS)*time.Second)

	if err != nil {
		return nil, fmt.Errorf("postgres: NewClient: %s", err)
	}

	if err := InitDB(ctx, pool, build); err != nil {
		return nil, fmt.Errorf("postgres: NewClient: %s", err)
	}

	return pool, nil
}
