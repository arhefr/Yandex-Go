package main

import (
	"calculator/config"
	"calculator/internal/agent/client"
	router "calculator/internal/transport/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup Logger
	file, err := os.OpenFile(".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error open logfile", err)
	}
	log.SetOutput(file)
	log.SetLevel(log.DebugLevel)

	// Setup Config
	config := config.NewConfig()

	// Run Agent's
	client.RunWorkers(config.AgentConfig, config.RouterConfig.Port)

	// Run Orchestrator
	router := router.NewRouter(config.RouterConfig)
	router.Run()
}
