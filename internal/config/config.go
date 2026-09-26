package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return "", errors.New("Error loading DB_URL")
	}

	return dbURL, nil
}
