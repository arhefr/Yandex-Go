package router

import "errors"

var (
	ErrIncorrectJSON = errors.New("error incorrect JSON")
	ErrIncorrectID   = errors.New("error incorrect ID")
	ErrNotFoundTask  = errors.New("error no available tasks")
)
