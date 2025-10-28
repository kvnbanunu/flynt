package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Uses bcrypt and cost from .env to hash password
func HashPassword(password string) (string, error) {
	pass := []byte(password)

	hashed, err := bcrypt.GenerateFromPassword(pass, CFG.Cost)
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
