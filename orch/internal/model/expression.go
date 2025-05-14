package model

import (
	"github.com/google/uuid"
)

type Expression struct {
	UserID string `json:"-"`
	ID     string `json:"id"`
	Status string `json:"status"`
	Expr   string `json:"expression"`
	Result string `json:"result"`
}

func NewExpression(userID string) *Expression {
	return &Expression{UserID: userID, ID: uuid.NewString(), Status: StatusWait, Expr: "", Result: ""}
}
