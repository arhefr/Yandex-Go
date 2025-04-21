package config

import (
	"os"

	router "github.com/arhefr/Yandex-Go/orch/internal/transport/http"
)

func NewRouterCfg() *router.Config {
	port := os.Getenv("PORT")

	return &router.Config{
		Port: port,
	}
}
