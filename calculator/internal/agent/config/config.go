package config

import "time"

type Config struct {
	Port string
	Host string
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
