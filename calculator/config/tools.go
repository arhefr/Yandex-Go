package config

import (
	"os"
	"strconv"
	"time"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvTime(key string, defaultVal string) time.Duration {
	if duration, err := time.ParseDuration(getEnv(key, defaultVal) + "ms"); err == nil {
		return time.Duration(duration)
	}
	return time.Millisecond * 10
}

func getEnvInt(key string, defaultVal string) int {
	if num, err := strconv.Atoi(getEnv(key, defaultVal)); err == nil {
		return num
	}
	return 1
}
