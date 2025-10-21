package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Profile struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Bio       string    `json:"bio" db:"bio"`
	ImgURL    string    `json:"img_url" db:"img_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// get user profile by their id
func (db *DB) GetProfileByUserID(userID int) (*Profile, error) {
	query := `
	SELECT id, user_id, name, email, bio, img_url, created_at, updated_at
	FROM profiles
	WHERE user_id = ?
	`
	var p Profile
	err := db.Get(&p, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Profile not found: %w", err)
		}
		return nil, fmt.Errorf("Failed to get profile: %w", err)
	}
	return &p, nil
}


// Update user profile
func (db *DB) UpdateProfile(userID int, updates map[string]any) (*Profile, error) {
	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	setParts := []string{}
	args := []any{}

	// Build dynamic query
	for col, val := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = ?", col))
		args = append(args, val)
	}

	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())

	args = append(args, userID)

	query := fmt.Sprintf(`
		UPDATE profiles
		SET %s
		WHERE user_id = ?
		RETURNING id, user_id, name, email, bio, img_url, created_at, updated_at
	`, strings.Join(setParts, ", "))

	var p Profile
	err := db.Get(&p, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to update profile: %w", err)
	}

	return &p, nil
}