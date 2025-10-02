package database

import (
	"flynt/internal/utils"
	"fmt"
)

type AccountLoginRequest struct {
	Email     string `json:"email"`
	Password string `json:"password"`
}

// check if login details match in database
func (db *DB) ValidateLogin(req AccountLoginRequest) (*User, error) {
	query := `SELECT * FROM user WHERE email = ?`
	var existingUser User
	err := db.Get(&existingUser, query, req.Email)
	if err != nil { // user does not exist, but do not expose info
		return nil, fmt.Errorf("Failed to login")
	}

	match, err := utils.CompareHash(existingUser.Password, req.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to hash password: %w", err)
	}
	
	if !match {
		return nil, fmt.Errorf("Failed to login")
	}

	existingUser.Password = ""
	
	return &existingUser, nil
}
