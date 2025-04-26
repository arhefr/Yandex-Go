package main

import (
	"os"

	"github.com/arhefr/Yandex-Go/orch/config"
	"github.com/arhefr/Yandex-Go/orch/internal/repository"
	"github.com/arhefr/Yandex-Go/orch/internal/service"
	router "github.com/arhefr/Yandex-Go/orch/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

func init() {
	file, err := os.OpenFile(".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error open logfile", err)
	}
	log.SetOutput(file)
	log.SetLevel(log.DebugLevel)
}

func main() {
	repo := repository.NewRepo()
	service := service.NewService(repo)
	handler := router.NewHandler(service)

	router := router.NewRouter(config.NewRouterCfg(), handler)
	router.Run()
}
