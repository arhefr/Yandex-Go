package main

import (
	"calculator/config"
	"calculator/internal/agent/client"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Setup Logger
	file, err := os.OpenFile(".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error open logfile", err)
	}
	log.SetOutput(file)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Setup Config
	cfg := config.NewAgentConfig()

	// Run Agent's
	client.RunWorkers(cfg)
}
