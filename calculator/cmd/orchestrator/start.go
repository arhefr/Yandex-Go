package main

import (
	"calculator/config"
	router "calculator/internal/transport/http"
)

func main() {
	config := config.NewConfig()

	router := router.NewRouter(config.RouterConfig)
	router.Run()
}
