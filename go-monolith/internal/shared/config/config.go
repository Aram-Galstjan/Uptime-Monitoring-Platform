package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppEnv            string
	LogLevel          string
	HTTPPort          string
	DatabaseURL       string
	SchedulerInterval time.Duration
	WorkerCount       int
}

func Load() (Config, error) {
	intervalRaw := getEnv("SCHEDULER_INTERVAL", "30s")
	interval, err := time.ParseDuration(intervalRaw)
	if err != nil {
		return Config{}, fmt.Errorf("parse SCHEDULER_INTERVAL: %w", err)
	}

	workerCount, err := strconv.Atoi(getEnv("SCHEDULER_WORKER_COUNT", "5"))
	if err != nil {
		return Config{}, fmt.Errorf("parse SCHEDULER_WORKER_COUNT: %w", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}

	return Config{
		AppEnv:            getEnv("APP_ENV", "development"),
		LogLevel:          getEnv("LOG_LEVEL", "info"),
		HTTPPort:          getEnv("SCHEDULER_HTTP_PORT", "8082"),
		DatabaseURL:       databaseURL,
		SchedulerInterval: interval,
		WorkerCount:       workerCount,
	}, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
