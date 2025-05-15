package repository

import (
	"context"
	"fmt"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryUsers struct {
	db postgres.Client
}

func NewRepositoryUsers(ctx context.Context, db *pgxpool.Pool) *RepositoryUsers {
	return &RepositoryUsers{db: db}
}

func (r *RepositoryUsers) SignIn(ctx context.Context, user *model.User) (err error) {

	if _, err := r.db.Exec(ctx, `
	INSERT INTO users (id, login, password)
    VALUES ($1, $2, $3)`, user.ID, user.Login, user.Password); err != nil {
		return fmt.Errorf("error internal server")
	}

	return nil
}

func (r *RepositoryUsers) ParseID(ctx context.Context, user *model.User) (*model.User, error) {

	row := r.db.QueryRow(ctx, `SELECT * FROM users WHERE login=$1 AND password=$2`, user.Login, user.Password)
	if err := row.Scan(&user.ID, &user.Login, &user.Password); err == pgx.ErrNoRows {
		return nil, fmt.Errorf("error internal server")
	}

	return user, nil
}

func (r *RepositoryUsers) Check(ctx context.Context, login string) bool {
	var tmp interface{}

	row := r.db.QueryRow(ctx, `SELECT * FROM users WHERE login=$1`, login)
	if err := row.Scan(&tmp); err != nil {
		fmt.Println("tmp", tmp)
		return false
	}

	fmt.Println("tmp", tmp)
	return true
}
