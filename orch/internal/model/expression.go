package model

import (
	repeatible "github.com/arhefr/Yandex-Go/orch/pkg/utils"
)

type Expression struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Expr   string `json:"expression"`
	Result string `json:"result"`
}

func NewExpression() *Expression {
	return &Expression{ID: repeatible.NewCryptoRand(1000000), Status: StatusWait, Expr: "", Result: ""}
}
