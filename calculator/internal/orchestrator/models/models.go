package models

import (
	"calculator/pkg/parser"
	"fmt"
)

type Response struct {
	ID      string `json:"id"`
	Status  int    `json:"status"`
	Request string `json:"request"`
	Result  string `json:"result"`
}

type Task struct {
	ID        int     `json:"id"`
	Operand1  float64 `json:"arg1"`
	Operand2  float64 `json:"arg2"`
	Operation string  `json:"operation"`
}

type Expression struct {
	ID      int    `json:"id"`
	Status  string `json:"status"`
	Result  string `json:"result"`
	Request `json:"-"`
	Parser  `json:"-"`
}

type Request struct {
	Expression string `json:"expression"`
}

type Parser struct {
	Nums []float64
	Ops  []parser.Operator
}

func NewExpression(id int, request Request) Expression {
	nums, ops, err := parser.GetNumsOps(request.Expression)
	if err != nil {
		return Expression{ID: id, Status: "done", Result: "incorrect math expression", Request: request}
	}

	return Expression{ID: id, Status: "processing", Request: request, Parser: Parser{nums, ops}}
}

func (e Expression) GetTask() *Task {
	op := e.Parser.Ops[0]
	num1, num2 := op.GetOperands(e.Parser.Nums)
	switch op.Name {
	case "+":
	}
	return &Task{e.ID, num1, num2, op.Name}
}

func (t Task) String() string {
	return fmt.Sprintf("ID: %d Task: %.3f%s%.3f", t.ID, t.Operand1, t.Operation, t.Operand2)
}
