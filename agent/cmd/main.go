package main

import (
	"github.com/arhefr/Yandex-Go/agent/config"
	"github.com/arhefr/Yandex-Go/agent/internal/client"
)

func main() {
	cfg := config.NewServiceCfg()

	client.RunWorkers(cfg)
}
