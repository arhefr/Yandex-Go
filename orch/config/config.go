package config

import (
	"os"

	router "github.com/arhefr/Yandex-Go/orch/internal/transport/http"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func NewRouterCfg() *router.Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Warn("error missing env")
	}

	port := os.Getenv("PORT")

	return &router.Config{
		Port: port,
	}
}
