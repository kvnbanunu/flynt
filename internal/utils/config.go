package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env     string
	Client  string
	DBPath  string
	Port    string
	Cost    int
	Secret  string
	Context string
	AdminID int
}

var CFG Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file: %w", err)
	}

	env := os.Getenv("ENVIRONMENT")
	client := os.Getenv("CLIENT_URL")
	path := os.Getenv("DB_PATH")
	port := os.Getenv("PORT")
	cost := os.Getenv("COST")
	secret := os.Getenv("JWT_SECRET")
	context := os.Getenv("JWT_CONTEXT")
	adminID := os.Getenv("ADMIN_ID")

	if !ValidateFields(env, client, path, port, cost, secret, context, adminID) {
		return fmt.Errorf("Missing environment variables")
	}

	// convert any values if needed
	costVal, err := strconv.Atoi(cost)
	if err != nil {
		return err
	}

	adminIDVal, err := strconv.Atoi(adminID)
	if err != nil {
		return err
	}

	CFG.Env = env
	CFG.Client = client
	CFG.DBPath = path
	CFG.Port = port
	CFG.Cost = costVal
	CFG.Secret = secret
	CFG.Context = context
	CFG.AdminID = adminIDVal

	return nil
}

func GetConfig() *Config {
	return &CFG
}
