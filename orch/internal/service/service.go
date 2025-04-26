package service

import (
	"github.com/arhefr/Yandex-Go/orch/internal/model"
)

type Repository interface {
	Get() []model.Request
	GetByID(id string) (model.Request, bool)
	Add(expr model.Expression) (string, error)
	Replace(id string, req model.Request)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Get() []model.Request {
	return s.repo.Get()
}

func (s *Service) GetByID(id string) (model.Request, bool) {
	return s.repo.GetByID(id)
}

func (s *Service) Add(expr model.Expression) (string, error) {
	return s.repo.Add(expr)
}

func (s *Service) Replace(id string, req model.Request) {
	s.repo.Replace(id, req)
}
