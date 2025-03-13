package config

import (
	"os"
	"strconv"
	"time"
)

func get(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getTime(key string, defaultVal string) time.Duration {
	if duration, err := time.ParseDuration(get(key, defaultVal) + "ms"); err == nil {
		return time.Duration(duration)
	}
	return time.Millisecond
}

func getInt(key string, defaultVal string) int {
	if num, err := strconv.Atoi(get(key, defaultVal)); err == nil {
		return num
	}
	return 1
}
