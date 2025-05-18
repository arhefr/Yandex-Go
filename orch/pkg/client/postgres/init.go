package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(ctx context.Context, pool *pgxpool.Pool, build string) error {
	if _, err := pool.Exec(ctx, build); err != nil {
		return fmt.Errorf("postgres: InitDB: %s", err)
	}

	return nil
}
