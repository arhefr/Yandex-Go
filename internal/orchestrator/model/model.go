package model

import (
	"fmt"

	"github.com/arhefr/Yandex-Go/pkg/parser"
)

const (
	StatusWait = "waiting"
	StatusCalc = "calculating"
	StatusDone = "done"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	ID      string `json:"id"`
	Status  int    `json:"status"`
	Request string `json:"request"`
	Result  string `json:"result"`
}

type Task struct {
	ID        string  `json:"id"`
	Operand1  float64 `json:"arg1"`
	Operand2  float64 `json:"arg2"`
	Operation string  `json:"operation"`
}

type Expression struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Result  string `json:"result"`
	*Parser `json:"-"`
}

type Parser struct {
	Nums []float64
	Ops  []parser.Operator
}

func NewExpression(id string, request *Request) Expression {
	nums, ops, err := parser.GetNumsOps(request.Expression)
	if err != nil {
		return Expression{ID: id, Status: StatusDone, Result: fmt.Sprintf("%s", err)}
	}

	return Expression{ID: id, Status: StatusWait, Parser: &Parser{nums, ops}}
}

func (e Expression) GetTask() *Task {
	op := e.Parser.Ops[0]
	num1, num2 := op.GetOperands(e.Parser.Nums)
	return &Task{e.ID, num1, num2, op.Name}
}
