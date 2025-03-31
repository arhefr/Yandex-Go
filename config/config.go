package config

import (
	"sync"

	"github.com/arhefr/Yandex-Go/internal/agent/service"
	router "github.com/arhefr/Yandex-Go/internal/orchestrator/transport/http"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

type Config struct {
	RouterConfig router.Config
	AgentConfig  service.Config
}

func NewRouterCfg() *router.Config {
	if err := godotenv.Load("config/enviroment.env"); err != nil {
		log.Fatal("error missing enviroment file")
	}

	return &router.Config{
		Port:     envStr("PORT", dfltPort),
		PathAdd:  envStr("PATH_ADD", dfltPathAdd),
		PathGet:  envStr("PATH_GET", dfltPathGet),
		PathTask: envStr("PATH_TASK", dfltPathTask),
	}
}

func NewServiceCfg() *service.Config {
	if err := godotenv.Load("config/enviroment.env"); err != nil {
		log.Fatal("error missing enviroment file")
	}

	return &service.Config{
		Port:             envStr("PORT", dfltPort),
		Path:             envStr("PATH_TASK", dfltPathTask),
		AgentsValue:      envInt("COMPUTING_POWER", dfltAgentTick),
		AgentPeriodicity: envTime("AGENT_PERIODICITY_MS", dfltAgentTick),
		OperTime: service.OperTime{
			Add: envTime("TIME_ADDITION_MS", dfltOperAdd),
			Sub: envTime("TIME_SUBTRACTION_MS", dfltOperSub),
			Mul: envTime("TIME_MULTIPLICATION_MS", dfltOperMul),
			Div: envTime("TIME_DIVISION_MS", dfltOperDiv)},

		WG: &sync.WaitGroup{},
	}
}
