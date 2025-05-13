package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var build = `
CREATE TABLE IF NOT EXISTS users 
(
id text,
login text,
password text
);

CREATE TABLE IF NOT EXISTS expressions 
(
id text,
status text,
expression text,
result text
);`

func InitDB(ctx context.Context, pool *pgxpool.Pool) error {
	if _, err := pool.Exec(ctx, build); err != nil {
		return fmt.Errorf("postgres: InitDB: %s", err)
	}

	return nil
}
