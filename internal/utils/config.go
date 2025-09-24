package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env    string
	DBPath string
	Port   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	env := os.Getenv("ENVIRONMENT")
	path := os.Getenv("DB_PATH")
	port := os.Getenv("PORT")

	if env == "" || path == "" || port == "" {
		return nil, fmt.Errorf("Missing environment variables")
	}
	
	config := Config{
		Env:    env,
		DBPath: path,
		Port:   port,
	}

	return &config, nil
}
