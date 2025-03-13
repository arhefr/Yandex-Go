package models

import (
	"time"
)

type Response struct {
	ID     string  `json:"id"`
	Result float64 `json:"result"`
}

type OperationTime struct {
	Add time.Duration
	Sub time.Duration
	Mul time.Duration
	Div time.Duration
}
