package errors

import (
	"errors"
)

var (
	ErrDivisionByZero = errors.New("error division by zero")
	ErrIncorrectExpr  = errors.New("error incorrect math expression")
)
