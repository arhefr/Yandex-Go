package main

import (
	"calculator/config"
	"calculator/internal/orchestrator/transport/http/router"
)

func main() {
	config := config.NewConfig()

	router := router.NewRouter(config.RouterConfig)
	router.Run()
}
