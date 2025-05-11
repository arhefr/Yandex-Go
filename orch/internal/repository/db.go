package repository

import (
	"context"
	"fmt"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

var build = `
CREATE TABLE IF NOT EXISTS expressions (
    id text,
	status text,
	expression text,
	result text
);
`

type Repository struct {
	db postgres.Client
}

func NewRepository(ctx context.Context, db *pgxpool.Pool) (*Repository, error) {
	if _, err := db.Exec(ctx, build); err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}

func (r *Repository) Get(ctx context.Context) (exprs []model.Expression, err error) {
	q := "SELECT * FROM expressions"

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expr model.Expression
		if err = rows.Scan(&expr.ID, &expr.Status, &expr.Expr, &expr.Result); err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)
	}

	return exprs, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (expr model.Expression, err error) {
	q := "SELECT * FROM expressions WHERE id=$1"

	row := r.db.QueryRow(ctx, q, id)
	if err := row.Scan(&expr.ID, &expr.Status, &expr.Expr, &expr.Result); err != nil {
		return model.Expression{}, fmt.Errorf("repository: GetByID: %s", err)
	}

	return expr, nil
}

func (r *Repository) Add(ctx context.Context, expr model.Expression) (err error) {
	q := `
		INSERT INTO expressions (id, status, expression, result)
        VALUES ($1, $2, $3, $4)`

	if _, err := r.db.Exec(ctx, q, expr.ID, expr.Status, expr.Expr, expr.Result); err != nil {
		return fmt.Errorf("repository: Add: %s", err)
	}

	return nil
}

func (r *Repository) Replace(ctx context.Context, id, status, result string) error {
	q := `UPDATE expressions SET result=$1, status=$2 WHERE id=$3`

	if _, err := r.db.Exec(ctx, q, result, model.StatusDone, id); err != nil {
		return fmt.Errorf("repository: Replace: %s", err)
	}

	return nil
}
