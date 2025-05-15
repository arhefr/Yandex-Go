package services

import (
	"context"
	"fmt"

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
	ParseID(ctx context.Context, user *model.User) (*model.User, error)
	Check(ctx context.Context, login string) bool
}

func NewServiceUsers(db RepositoryUsers, tm *jwt.Manager, ph *hash.Hasher) *ServiceUsers {
	return &ServiceUsers{db: db, tm: tm, ph: ph}
}

func (su *ServiceUsers) SignIn(ctx context.Context, user *model.User) (err error) {
	user.Password = su.ph.Hash(user.Password)
	fmt.Println(user)
	if exists := su.db.Check(ctx, user.Login); exists {
		return fmt.Errorf("error login already exists")
	}
	return su.db.SignIn(ctx, user)
}

func (su *ServiceUsers) LogIn(ctx context.Context, user *model.User) (jwt string, err error) {
	user.Password = su.ph.Hash(user.Password)
	if user, err := su.db.ParseID(ctx, user); err != nil {

		token, err := su.tm.NewJWT(user.Login, user.ID)
		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", fmt.Errorf("services: LogIn: %s", "error inccorrect data")
}

func (su *ServiceUsers) ParseJWT(jwt string) (user *model.User, err error) {
	claims, err := su.tm.Parse(jwt)
	if err != nil {
		return nil, err
	}

	return &model.User{ID: claims["id"].(string), Login: claims["login"].(string)}, nil
}
