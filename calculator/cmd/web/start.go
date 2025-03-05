package main

import (
	"calculator/config"
	agent "calculator/internal/agent/client"
	"calculator/internal/orchestrator/transport/http/router"
)

func main() {
	config := config.NewConfig()

	router := router.NewRouter(config.RouterConfig)
	agent.RunWorkers(config.AgentConfig)
	router.Run()
}
