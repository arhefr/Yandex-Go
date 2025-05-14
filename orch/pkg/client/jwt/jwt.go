package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	NewJWT(userID string, login string) (string, error)
	Parse(token string) (string, error)
}

type Manager struct {
	key string
}

func NewManager(key string) *Manager {
	return &Manager{key: key}
}

func (m *Manager) NewJWT(login string, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"login": login,
	})
	return token.SignedString([]byte(m.key))
}

func (m *Manager) Parse(accToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("jwt: Parse: error unexpected signing method")
		}

		return []byte(m.key), nil
	})

	if err != nil {
		return nil, fmt.Errorf("jwt: Parse: %s", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("jwt: Parse: error get claims from token")
	}

	return claims, nil
}
