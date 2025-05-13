package services

import (
	repository "github.com/arhefr/Yandex-Go/orch/internal/repositories"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/hash"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/jwt"
)

type Services struct {
	*ServiceExpressions
	*ServiceUsers
}

type Deps struct {
	Repos          *repository.Repositories
	SafeMap        TempRepo
	TokenManager   *jwt.Manager
	PasswordHasher *hash.Hasher
}

func NewServices(deps Deps) *Services {
	return &Services{
		NewServiceExpressions(deps.Repos.RepositoryExpressions, deps.SafeMap),
		NewServiceUsers(deps.Repos.RepositoryUsers),
	}
}
