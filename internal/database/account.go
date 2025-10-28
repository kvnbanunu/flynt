package database

import (
	"fmt"

	"flynt/internal/utils"
)

type AccountLoginRequest struct {
	LoginType string `json:"type"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password"`
}

// check if login details match in database
func (db *DB) ValidateLogin(req AccountLoginRequest) (*User, error) {
	query := fmt.Sprintf("SELECT * FROM user WHERE %s = ? COLLATE NOCASE", req.LoginType)

	qparam := req.Email
	if req.LoginType == "username" {
		qparam = req.Username
	}
	
	var existingUser User
	err := db.Get(&existingUser, query, qparam)
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
