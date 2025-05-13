package services

import (
	"context"
)

type ServiceUsers struct {
	db RepositoryUsers
}

type RepositoryUsers interface {
	SignIn(ctx context.Context, login, password string) (err error)
	LogIn(ctx context.Context, login, password string) (jwt string, err error)
}

func NewServiceUsers(db RepositoryUsers) *ServiceUsers {
	return &ServiceUsers{db: db}
}

func (su *ServiceUsers) SignIn(ctx context.Context, login, password string) (err error) {
	return su.db.SignIn(ctx, login, password)
}

func (su *ServiceUsers) LogIn(ctx context.Context, login, password string) (jwt string, err error) {
	return su.db.LogIn(ctx, login, password)
}
