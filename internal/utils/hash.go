package utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// Uses bcrypt and cost from .env to hash password
func HashPassword(password string) (string, error) {
	pass := []byte(password)
	cost := os.Getenv("COST")
	costVal, err := strconv.Atoi(cost)
	if err != nil {
		costVal = bcrypt.DefaultCost
	}

	hashed, err := bcrypt.GenerateFromPassword(pass, costVal)
	if err != nil {
		return "", err
	}
	
	return string(hashed), nil
}

// Compare hash from db to raw string
func CompareHash(hashed, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
