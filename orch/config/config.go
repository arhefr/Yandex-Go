package config

import (
	"fmt"
	"os"
	"strconv"

	router "github.com/arhefr/Yandex-Go/orch/internal/transport/http/router"
	"github.com/arhefr/Yandex-Go/orch/pkg/client/postgres"
)

type Config struct {
	API router.Config
	DB  postgres.DBConfig
	Storage
}

func NewConfig() (*Config, error) {
	routerPort, dbHost, dbPort, dbUser, dbPassword, dbName, dbMaxAtmps, dbDelayAtmps, jwtKey, hashSalt :=
		os.Getenv("PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_MAX_ATMPS"),
		os.Getenv("DB_DELAY_ATMPS_S"),
		os.Getenv("JWT_KEY"),
		os.Getenv("HASH_SALT")

	if routerPort == "" || dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbMaxAtmps == "" || dbDelayAtmps == "" || jwtKey == "" || hashSalt == "" {
		return &Config{}, fmt.Errorf("config: NewConfig: error missing enviroment params")
	}

	dbMaxAtmpsInt, err1 := strconv.Atoi(dbMaxAtmps)
	dbDelayAtmpsInt, err2 := strconv.Atoi(dbDelayAtmps)
	if err1 != nil || err2 != nil {
		return &Config{}, fmt.Errorf("config: NewConfig: error wrong enviroment param")
	}

	return &Config{
		API: router.Config{
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
		Storage: Storage{
			JWTkey:   jwtKey,
			HashSalt: hashSalt,
		},
	}, nil
}
