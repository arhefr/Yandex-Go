package main

import (
	"calculator/config"
	"calculator/internal/agent/client"
	router "calculator/internal/transport/http"
)

func main() {
	config := config.NewConfig()

	router := router.NewRouter(config.RouterConfig)
	client.RunWorkers(config.AgentConfig, config.RouterConfig.Port)
	router.Run()
}
