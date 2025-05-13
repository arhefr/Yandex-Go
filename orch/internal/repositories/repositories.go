package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	*RepositoryExpressions
	*RepositoryUsers
}

func NewRepositories(ctx context.Context, pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		NewRepositoryExpressions(ctx, pool),
		NewRepositoryUsers(ctx, pool),
	}
}
