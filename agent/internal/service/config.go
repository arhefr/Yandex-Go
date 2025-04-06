package service

import (
	"sync"
	"time"
)

type Config struct {
	Port string

	AgentsValue      int
	AgentPeriodicity time.Duration

	WG *sync.WaitGroup
}
