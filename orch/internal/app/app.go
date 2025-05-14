package app

import (
	"context"
	"os"

	"github.com/arhefr/Yandex-Go/orch/config"
	repository "github.com/arhefr/Yandex-Go/orch/internal/repositories"
	"github.com/arhefr/Yandex-Go/orch/internal/services"
	"github.com/arhefr/Yandex-Go/orch/internal/transport/http/handlers"
	router "github.com/arhefr/Yandex-Go/orch/internal/transport/http/router"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/hash"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/jwt"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
	"github.com/arhefr/Yandex-Go/orch/pkg/logger"
	log "github.com/sirupsen/logrus"
)

func Run() {
	logfile, err := os.OpenFile(".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	logger := logger.NewLogger(logfile, log.TraceLevel)

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}

	db, err := postgres.NewClient(logger, context.TODO(), cfg.DB)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	postgres.InitDB(context.TODO(), db)

	tokenManager := jwt.NewManager(cfg.Storage.JWTkey)
	passwordHasher := hash.NewHasher(cfg.Storage.HashSalt)
	repos := repository.NewRepositories(context.TODO(), db)
	deps := services.Deps{
		Repos:          repos,
		SafeMap:        repository.NewSafeMap(),
		TokenManager:   tokenManager,
		PasswordHasher: passwordHasher,
	}
	services := services.NewServices(deps)
	handler := handlers.NewHandler(services)
	router := router.NewRouter(&cfg.API, &handler)
	router.Run()
}
