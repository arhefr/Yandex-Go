package repository

import (
	"context"
	"fmt"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryExpressions struct {
	db postgres.Client
}

func NewRepositoryExpressions(ctx context.Context, db *pgxpool.Pool) *RepositoryExpressions {
	return &RepositoryExpressions{db: db}
}

func (r *RepositoryExpressions) Get(ctx context.Context, userID string) (exprs []model.Expression, err error) {

	rows, err := r.db.Query(ctx, `SELECT * FROM expressions WHERE userID=$1`, userID)
	if err != nil {
		return nil, fmt.Errorf("repository: Get: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var expr model.Expression
		if err = rows.Scan(&expr.UserID, &expr.ID, &expr.Status, &expr.Expr, &expr.Result); err != nil {
			return nil, fmt.Errorf("repository: Get: %s", err)
		}
		exprs = append(exprs, expr)
	}

	return exprs, nil
}

func (r *RepositoryExpressions) GetByID(ctx context.Context, userID string, id string) (expr model.Expression, err error) {

	row := r.db.QueryRow(ctx, `SELECT * FROM expressions WHERE userID=$1 AND id=$2`, userID, id)
	if err := row.Scan(&expr.UserID, &expr.ID, &expr.Status, &expr.Expr, &expr.Result); err != nil {
		return model.Expression{}, fmt.Errorf("repository: GetByID: %s", err)
	}

	return expr, nil
}

func (r *RepositoryExpressions) Add(ctx context.Context, expr model.Expression) (err error) {

	if _, err := r.db.Exec(ctx, `
	INSERT INTO expressions (userID, id, status, expression, result)
    VALUES ($1, $2, $3, $4, $5)`, expr.UserID, expr.ID, expr.Status, expr.Expr, expr.Result); err != nil {
		return fmt.Errorf("repository: Add: %s", err)
	}

	return nil
}

func (r *RepositoryExpressions) Replace(ctx context.Context, id, status, result string) error {

	if _, err := r.db.Exec(ctx, `UPDATE expressions SET result=$1, status=$2 WHERE id=$3`, result, status, id); err != nil {
		return fmt.Errorf("repository: Replace: %s", err)
	}

	return nil
}
