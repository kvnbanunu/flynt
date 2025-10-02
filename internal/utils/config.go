package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env    string
	DBPath string
	Port   string
	Cost   int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	env := os.Getenv("ENVIRONMENT")
	path := os.Getenv("DB_PATH")
	port := os.Getenv("PORT")
	cost := os.Getenv("COST")

	if env == "" || path == "" || port == "" || cost == "" {
		return nil, fmt.Errorf("Missing environment variables")
	}

	// convert any values if needed
	costVal, err := strconv.Atoi(cost)
	if err != nil {
		return nil, err
	}

	config := Config{
		Env:    env,
		DBPath: path,
		Port:   port,
		Cost:   costVal,
	}

	return &config, nil
}
