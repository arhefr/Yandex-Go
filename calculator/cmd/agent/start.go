package main

import (
	"calculator/config"
	"calculator/internal/agent/client"
)

func main() {
	config := config.NewConfig()

	client.RunWorkers(config.AgentConfig, config.RouterConfig.Host)
}
