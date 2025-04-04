package config

import (
	"os"
	"strconv"
	"time"
)

// funcs for parse params from enviroment file

func envStr(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func envTime(key string, defaultVal string) time.Duration {
	if duration, err := time.ParseDuration(envStr(key, defaultVal) + "ms"); err == nil {
		return time.Duration(duration)
	}

	panic("error enviroment data")
}

func envInt(key string, defaultVal string) int {
	if num, err := strconv.Atoi(envStr(key, defaultVal)); err == nil {
		return num
	}

	panic("error enviroment data")
}
