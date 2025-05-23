package config

import (
	"os"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Host string
	Port string

	AgentsValue      int
	AgentPeriodicity time.Duration

	WG *sync.WaitGroup
}

func NewServiceCfg() *Config {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	agentsValue, err := strconv.Atoi(os.Getenv("COMPUTING_POWER"))
	if err != nil {
		log.Fatal("error incorrect enviroment data:", os.Getenv("COMPUTING_POWER"))
	}
	agentPeriodicity, err := time.ParseDuration(os.Getenv("AGENT_PERIODICITY_MS") + "ms")
	if err != nil {
		log.Fatal("error incorrect enviroment data:", os.Getenv("AGENT_PERIODICITY_MS"))
	}

	return &Config{
		Port: port,
		Host: host,

		AgentsValue:      agentsValue,
		AgentPeriodicity: agentPeriodicity,

		WG: &sync.WaitGroup{},
	}
}
