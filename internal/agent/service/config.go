package service

import (
	"sync"
	"time"
)

type Config struct {
	Port string
	Path string

	AgentsValue      int
	AgentPeriodicity time.Duration

	WG *sync.WaitGroup

	OperTime
}

type OperTime struct {
	Add time.Duration
	Sub time.Duration
	Mul time.Duration
	Div time.Duration
}
