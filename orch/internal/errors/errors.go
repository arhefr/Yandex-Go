package errors

import (
	"errors"
)

var (
	DivisionByZero = errors.New("error division by zero")
	IncorrectExpr  = errors.New("error incorrect math expression")

	IncorrectJSON = errors.New("error incorrect JSON")
	IncorrectID   = errors.New("error incorrect ID")
	NotFoundTask  = errors.New("error no available tasks")

	InvalidData   = errors.New("error wrong data")
	NotAuthorized = errors.New("error not authorized")

	CannotConnect = errors.New("error cannot connect to server")

	Common = errors.New("error something went wrong")
)
