package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadAPIKey() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY not set in .env")
	}
	return apiKey, nil
}
