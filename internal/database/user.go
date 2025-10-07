package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"flynt/internal/utils"
)

// request payload for creating user
type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// request payload for updating user
type UpdateUserRequest struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	ImgURL   string `json:"img_url"`
	Bio      string `json:"bio"`
}

// query function to create new user in db
func (db *DB) CreateUser(req CreateUserRequest) (*User, error) {
	query := `
	INSERT INTO user (name, password, email)
	VALUES (?, ?, ?)
	RETURNING id, name, email, created_at, updated_at
	`

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to hash password: %w", err)
	}

	var user User
	err = db.Get(&user, query, req.Name, hashed, req.Email)
	if err != nil {
		return nil, fmt.Errorf("Failed to create user: %w", err)
	}

	return &user, nil
}

// query function to get user from db
func (db *DB) GetUserByID(id int) (*User, error) {
	query := `
	SELECT id, name, email, img_url, bio, created_at, updated_at
	FROM user
	WHERE id = ?
	`

	var user User
	err := db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found: %w", err)
		}
		return nil, fmt.Errorf("Failed to get user: %w", err)
	}

	return &user, nil
}

// query function to fetch user matching email
func (db *DB) GetUserByEmail(email string) (*User, error) {
	query := `
	SELECT id, name, email, img_url, bio, created_at, updated_at
	FROM user
	WHERE email = ?
	`

	var user User
	err := db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found: %w", err)
		}
		return nil, fmt.Errorf("Failed to get user: %w", err)
	}

	return &user, nil
}

// fetch all users from db
func (db *DB) GetAllUsers() ([]User, error) {
	query := `
	SELECT id, name, email, img_url, bio, created_at, updated_at
	FROM user
	ORDER BY name ASC
	`

	var users []User
	err := db.Select(&users, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get users: %w", err)
	}
	// Select will not error if no rows were found, the array will just be empty

	return users, nil
}

// update user row in database
func (db *DB) UpdateUser(id int, req UpdateUserRequest) (*User, error) {
	existingUser, err := db.GetUserByID(id) // check if user exists first
	if err != nil {
		return nil, err
	}

	// need to dynamically build query based on parts provided
	setParts := []string{}
	args := []any{}

	if req.Name != "" {
		setParts = append(setParts, "name = ?")
		args = append(args, req.Name)
	}

	if req.Password != "" {
		hashed, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, fmt.Errorf("Failed to hash password: %w", err)
		}
		setParts = append(setParts, "password = ?")
		args = append(args, hashed)
	}

	if req.Email != "" {
		setParts = append(setParts, "email = ?")
		args = append(args, req.Email)
	}

	if req.ImgURL != "" {
		setParts = append(setParts, "img_url = ?")
		args = append(args, req.ImgURL)
	}

	if req.Bio != "" {
		setParts = append(setParts, "bio = ?")
		args = append(args, req.Bio)
	}

	if len(setParts) == 0 { // no parts changed
		return existingUser, nil
	}

	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())

	// arg for WHERE id = ?
	args = append(args, id)

	query := fmt.Sprintf(`
		UPDATE user
		SET %s
		WHERE id = ?
		RETURNING id, name, email, img_url, bio, created_at, updated_at
	`, strings.Join(setParts, ", "))

	var user User
	err = db.Get(&user, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to update user: %w", err)
	}

	return &user, nil
}

// deletes user from db
func (db *DB) DeleteUser(id int) error {
	_, err := db.GetUserByID(id) // check if user exists
	if err != nil {
		return err
	}

	query := `DELETE FROM user WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No user deleted")
	}

	return nil
}
