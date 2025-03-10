package config

import (
	"calculator/internal/agent/models"
	router "calculator/internal/transport/http"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	RouterConfig router.Config
	AgentConfig  models.Config
}

func NewConfig() *Config {
	if err := godotenv.Load("config/enviroment.env"); err != nil {
		log.Fatal("error missing enviroment file")
	}

	return &Config{
		RouterConfig: router.Config{
			Port:     getEnv("PORT", "8080"),
			PathAdd:  getEnv("PATH_ADD", "/api/v1/calculate"),
			PathGet:  getEnv("PATH_GET", "/api/v1/expressions"),
			PathTask: getEnv("PATH_TASK", "/internal/task"),
		},
		AgentConfig: models.Config{
			Path:             getEnv("PATH_TASK", "/internal/task"),
			AgentsValue:      getEnvInt("COMPUTING_POWER", "10"),
			AgentPeriodicity: getEnvTime("AGENT_PERIODICITY_MS", "100"),
			OperationTime: models.OperationTime{
				Add: getEnvTime("TIME_ADDITION_MS", "100"),
				Sub: getEnvTime("TIME_SUBTRACTION_MS", "100"),
				Mul: getEnvTime("TIME_MULTIPLICATION_MS", "500"),
				Div: getEnvTime("TIME_DIVISION_MS", "500")},
		},
	}
}
