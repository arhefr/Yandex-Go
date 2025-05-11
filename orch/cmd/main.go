package main

import (
	"context"
	"os"

	"github.com/arhefr/Yandex-Go/orch/config"
	"github.com/arhefr/Yandex-Go/orch/internal/logger"
	"github.com/arhefr/Yandex-Go/orch/internal/repository"
	"github.com/arhefr/Yandex-Go/orch/internal/service"
	router "github.com/arhefr/Yandex-Go/orch/internal/transport/http"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {

	logfile, err := os.OpenFile(".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	logger := logger.NewLogger(logfile, log.TraceLevel)

	cfg, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := postgres.NewClient(logger, context.TODO(), cfg.DB)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	repo, err := repository.NewRepository(context.TODO(), db)
	if err != nil {
		logger.Fatal(err)
	}

	ts := repository.NewSafeMap()
	service := service.NewService(repo)
	handler := router.NewHandler(service, ts)
	router := router.NewRouter(&cfg.RouterConfig, &handler)
	router.Run()
}
