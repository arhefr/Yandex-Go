package config

import (
	"time"

	"github.com/arhefr/Yandex-Go/internal/agent/model"
	router "github.com/arhefr/Yandex-Go/internal/orchestrator/transport/http"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

type Config struct {
	RouterConfig router.Config
	AgentConfig
}

type AgentConfig struct {
	Port string
	Path string

	AgentsValue      int
	AgentPeriodicity time.Duration

	model.OperationTime
}

func NewConfig() *Config {
	if err := godotenv.Load("config/enviroment.env"); err != nil {
		log.Fatal("error missing enviroment file")
	}

	return &Config{
		RouterConfig: *NewRouterConfig(),
		AgentConfig:  *NewAgentConfig(),
	}
}

func NewRouterConfig() *router.Config {
	if err := godotenv.Load("config/enviroment.env"); err != nil {
		log.Fatal("error missing enviroment file")
	}

	return &router.Config{
		Port:     get("PORT", "8080"),
		PathAdd:  get("PATH_ADD", "/api/v1/calculate"),
		PathGet:  get("PATH_GET", "/api/v1/expressions"),
		PathTask: get("PATH_TASK", "/internal/task"),
	}
}

func NewAgentConfig() *AgentConfig {
	if err := godotenv.Load("config/enviroment.env"); err != nil {
		log.Fatal("error missing enviroment file")
	}

	return &AgentConfig{
		Port:             get("PORT", "8080"),
		Path:             get("PATH_TASK", "/internal/task"),
		AgentsValue:      getInt("COMPUTING_POWER", "10"),
		AgentPeriodicity: getTime("AGENT_PERIODICITY_MS", "100"),
		OperationTime: model.OperationTime{
			Add: getTime("TIME_ADDITION_MS", "100"),
			Sub: getTime("TIME_SUBTRACTION_MS", "100"),
			Mul: getTime("TIME_MULTIPLICATION_MS", "500"),
			Div: getTime("TIME_DIVISION_MS", "500")},
	}
}
