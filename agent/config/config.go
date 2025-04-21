package config

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/arhefr/Yandex-Go/agent/internal/service"
	log "github.com/sirupsen/logrus"
)

func NewServiceCfg() *service.Config {

	port := os.Getenv("PORT")
	agentsValue, err := strconv.Atoi(os.Getenv("COMPUTING_POWER"))
	if err != nil {
		log.Fatal("error incorrect enviroment data:", os.Getenv("COMPUTING_POWER"))
	}
	agentPeriodicity, err := time.ParseDuration(os.Getenv("AGENT_PERIODICITY_MS") + "ms")
	if err != nil {
		log.Fatal("error incorrect enviroment data:", os.Getenv("AGENT_PERIODICITY_MS"))
	}

	return &service.Config{
		Port: port,

		AgentsValue:      agentsValue,
		AgentPeriodicity: agentPeriodicity,

		WG: &sync.WaitGroup{},
	}
}
