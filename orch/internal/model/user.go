package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewUser() *User {
	return &User{ID: uuid.NewString()}
}
