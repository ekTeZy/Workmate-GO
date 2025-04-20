package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, продолжаем с переменными окружения")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config := &Config{
		Port: port,
	}

	return config
}
