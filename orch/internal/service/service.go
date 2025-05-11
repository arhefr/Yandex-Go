package service

import (
	"context"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
)

type Repository interface {
	Postgresql
}

type Postgresql interface {
	Get(ctx context.Context) ([]model.Expression, error)
	GetByID(ctx context.Context, id string) (model.Expression, error)
	Add(ctx context.Context, expr model.Expression) error
	Replace(ctx context.Context, id, status, result string) error
}

type Service struct {
	db Repository
}

func NewService(db Repository) *Service {
	return &Service{db: db}
}

func (s *Service) Get(ctx context.Context) ([]model.Expression, error) {
	return s.db.Get(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (model.Expression, error) {
	return s.db.GetByID(ctx, id)
}

func (s *Service) Add(ctx context.Context, expr model.Expression) (err error) {
	return s.db.Add(ctx, expr)
}

func (s *Service) Replace(ctx context.Context, id, status, result string) error {
	return s.db.Replace(ctx, id, status, result)
}
