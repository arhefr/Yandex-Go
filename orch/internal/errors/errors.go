package errors

import (
	"errors"
)

var (
	BadAuth            = errors.New("error bad login or password")
	LoginAlreadyExists = errors.New("error login already exists")
	IncorrectAuth      = errors.New("error incorrect login or password")
	IncorrectID        = errors.New("error incorrect id expr")
	NotAuthorized      = errors.New("error not authorized")
	NotFoundTask       = errors.New("error not found task")
	DivisionByZero     = errors.New("error division by zero")
	IncorrectExpr      = errors.New("error incorrect expr")
	IncorrectJSON      = errors.New("error incorrect JSON")
	InternalServer     = errors.New("error internal server")
)
