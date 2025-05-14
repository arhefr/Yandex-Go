package services

import (
	"context"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
)

type ServiceExpressions struct {
	db RepositoryExpressions
	tr TempRepo
}

type RepositoryExpressions interface {
	Get(ctx context.Context, id string) ([]model.Expression, error)
	GetByID(ctx context.Context, userID string, id string) (model.Expression, error)
	Add(ctx context.Context, expr model.Expression) error
	Replace(ctx context.Context, id, status, result string) error
}

type TempRepo interface {
	Add(req model.Request, key string)
	Replace(key string, req model.Request)
	Get() (arr []model.Request)
	GetByID(key string) (*model.Request, bool)
	Delete(key string)
}

func NewServiceExpressions(db RepositoryExpressions, tr TempRepo) *ServiceExpressions {
	return &ServiceExpressions{db: db, tr: tr}
}

func (se *ServiceExpressions) GetExprs(ctx context.Context, userID string) ([]model.Expression, error) {
	return se.db.Get(ctx, userID)
}

func (se *ServiceExpressions) GetExprByID(ctx context.Context, userID string, id string) (model.Expression, error) {
	return se.db.GetByID(ctx, userID, id)
}

func (se *ServiceExpressions) AddExpr(ctx context.Context, expr model.Expression) (err error) {
	return se.db.Add(ctx, expr)
}

func (se *ServiceExpressions) ReplaceExpr(ctx context.Context, id, status, result string) error {
	return se.db.Replace(ctx, id, status, result)
}

func (se *ServiceExpressions) AddReq(key string, req model.Request) {
	se.tr.Add(req, key)
}

func (se *ServiceExpressions) ReplaceReq(key string, req model.Request) {
	se.tr.Replace(key, req)
}

func (se *ServiceExpressions) GetReq() (arr []model.Request) {
	return se.tr.Get()
}

func (se *ServiceExpressions) GetReqByID(key string) (*model.Request, bool) {
	return se.tr.GetByID(key)
}

func (se *ServiceExpressions) DeleteReq(key string) {
	se.tr.Delete(key)
}
