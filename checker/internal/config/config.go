package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	WorkerCount int

	RequestTimeout int
	CheckInterval  int

	TelegramToken string

	EmailHost     string
	EmailPort     string
	EmailUser     string
	EmailPassword string
	EmailFrom     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить .env")
	}

	workerCount, _ := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	requestTimeout, _ := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT"))
	checkInterval, _ := strconv.Atoi(os.Getenv("CHECK_INTERVAL"))

	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		WorkerCount: workerCount,

		RequestTimeout: requestTimeout,
		CheckInterval:  checkInterval,

		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),

		EmailHost:     os.Getenv("EMAIL_HOST"),
		EmailPort:     os.Getenv("EMAIL_PORT"),
		EmailUser:     os.Getenv("EMAIL_USER"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
		EmailFrom:     os.Getenv("EMAIL_FROM"),
	}
}
