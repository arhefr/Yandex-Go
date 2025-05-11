package main

import (
	"github.com/arhefr/Yandex-Go/agent/config"
	"github.com/arhefr/Yandex-Go/agent/internal/client"
)

func main() {
	client.RunWorkers(config.NewServiceCfg())
}
