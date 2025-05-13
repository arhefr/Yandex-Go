package repository

import (
	"context"

	//"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryUsers struct {
	db postgres.Client
}

func NewRepositoryUsers(ctx context.Context, db *pgxpool.Pool) *RepositoryUsers {
	return &RepositoryUsers{db: db}
}

func (r *RepositoryUsers) SignIn(ctx context.Context, login, password string) (err error) {
	return nil
}

func (r *RepositoryUsers) LogIn(ctx context.Context, login, password string) (jwt string, err error) {
	return "", nil
}
