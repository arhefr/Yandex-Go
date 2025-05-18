package services

import (
	"context"

	"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/hash"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/jwt"
)

type ServiceUsers struct {
	db RepositoryUsers
	tm *jwt.Manager
	ph *hash.Hasher
}

type RepositoryUsers interface {
	SignIn(ctx context.Context, user *model.User) (err error)
	GetUserID(ctx context.Context, user *model.User) (string, error)
	Exists(ctx context.Context, user *model.User) (bool, error)
}

func NewServiceUsers(db RepositoryUsers, tm *jwt.Manager, ph *hash.Hasher) *ServiceUsers {
	return &ServiceUsers{db: db, tm: tm, ph: ph}
}

func (su *ServiceUsers) Exists(ctx context.Context, user *model.User) (bool, error) {
	return su.db.Exists(ctx, user)
}

func (su *ServiceUsers) CheckArgs(user *model.User) bool {
	if len(user.Login) < 3 || len(user.Password) < 8 {
		return false
	}

	return true
}

func (su *ServiceUsers) SignIn(ctx context.Context, user *model.User) (err error) {
	user.Password = su.ph.Hash(user.Password)
	return su.db.SignIn(ctx, user)
}

func (su *ServiceUsers) GetUserID(ctx context.Context, user *model.User) (string, error) {
	user.Password = su.ph.Hash(user.Password)
	return su.db.GetUserID(ctx, user)
}

func (su *ServiceUsers) GetJWT(uuid string) (string, error) {
	return su.tm.NewJWT(uuid)
}

func (su *ServiceUsers) ParseJWT(jwt string) (user *model.User, err error) {
	claims, err := su.tm.Parse(jwt)
	if err != nil {
		return nil, err
	}

	return &model.User{ID: claims["uuid"].(string)}, nil
}
