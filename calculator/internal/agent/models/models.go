package models

import (
	"time"
)

type Response struct {
	ID     int     `json:"id"`
	Result float64 `json:"result"`
}

type Config struct {
	Path string

	AgentsValue      int
	AgentPeriodicity time.Duration

	OperationTime
}

type OperationTime struct {
	Add time.Duration
	Sub time.Duration
	Mul time.Duration
	Div time.Duration
}
