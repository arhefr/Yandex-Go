package repository

import (
	"context"
	"fmt"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
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
		return fmt.Errorf("repository: SignIn: %s", err)
	}

	return nil
}

func (r *RepositoryUsers) GetUserID(ctx context.Context, user *model.User) (string, error) {
	var exists bool

	row := r.db.QueryRow(ctx, `SELECT EXISTS(SELECT * FROM users WHERE login=$1 AND password=$2)`, user.Login, user.Password)
	if err := row.Scan(&exists); err != nil {
		return "", fmt.Errorf("repository: GetUserID: %s", err)
	}

	if !exists {
		return "", fmt.Errorf("repository: GetUserID: %s", "wrong password")
	}

	row = r.db.QueryRow(ctx, `SELECT * FROM users WHERE login=$1 AND password=$2`, user.Login, user.Password)
	if err := row.Scan(&user.ID, &user.Login, &user.Password); err != nil {
		return "", fmt.Errorf("repository: GetUserID: %s", err)
	}

	return user.ID, nil
}

func (r *RepositoryUsers) Exists(ctx context.Context, user *model.User) (exists bool, err error) {

	row := r.db.QueryRow(ctx, `SELECT EXISTS(SELECT * FROM users WHERE login=$1)`, user.Login)
	if err := row.Scan(&exists); err != nil {
		return exists, fmt.Errorf("repository: Exists: %s", err)
	}

	return exists, nil
}
