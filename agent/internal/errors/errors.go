package errors

import (
	"errors"
)

var (
	IncorrectJSON = errors.New("error incorrect JSON")
	NotFoundTask  = errors.New("error no available tasks")
	CannotConnect = errors.New("error cannot connect to server")
)
