package config

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/arhefr/Yandex-Go/orch/internal/logger"
	router "github.com/arhefr/Yandex-Go/orch/internal/transport/http"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
)

type Config struct {
	RouterConfig router.Config
	DB           postgres.DBConfig
}

func NewConfig(logger *log.Logger) (*Config, error) {
	routerPort, dbHost, dbPort, dbUser, dbPassword, dbName, dbMaxAtmps, dbDelayAtmps :=
		os.Getenv("PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_MAX_ATMPS"),
		os.Getenv("DB_DELAY_ATMPS_S")

	logger.Info("Enviroment data: ", routerPort, dbHost, dbPort, dbUser, dbPassword, dbName, dbMaxAtmps, dbDelayAtmps)

	if routerPort == "" || dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbMaxAtmps == "" || dbDelayAtmps == "" {
		return &Config{}, fmt.Errorf("config: NewConfig: error missing enviroment params")
	}

	dbMaxAtmpsInt, err1 := strconv.Atoi(dbMaxAtmps)
	dbDelayAtmpsInt, err2 := strconv.Atoi(dbDelayAtmps)
	if err1 != nil || err2 != nil {
		return &Config{}, fmt.Errorf("config: NewConfig: error wrong enviroment param")
	}

	return &Config{
		RouterConfig: router.Config{
			Port: routerPort,
		},
		DB: postgres.DBConfig{
			Host:        dbHost,
			Port:        dbPort,
			User:        dbUser,
			Password:    dbPassword,
			Database:    dbName,
			MaxAtmps:    dbMaxAtmpsInt,
			DelayAtmpsS: dbDelayAtmpsInt,
		},
	}, nil
}
