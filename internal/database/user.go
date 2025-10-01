package database

import (
	"database/sql"
	"flynt/internal/utils"
	"fmt"
	"strings"
	"time"
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
	err = db.QueryRow(query, req.Name, hashed, req.Email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to create user: %w", err)
	}

	return &user, nil
}

// query function to get user from db
func (db *DB) GetUserByID(id int) (*User, error) {
	query := `
	SELECT id, name, email, created_at, updated_at
	FROM user
	WHERE id = ?
	`

	var user User
	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with id %d not found", id)
		}
		return nil, fmt.Errorf("Failed to get user: %w", err)
	}

	return &user, nil
}

// fetch all users from db
func (db *DB) GetAllUsers() ([]User, error) {
	query := `
	SELECT id, name, email, created_at, updated_at
	FROM user
	ORDER BY created_at DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to query users: %w", err)
	}
	defer rows.Close() // need to close or else read stalls

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating users: %w", err)
	}

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
		RETURNING id, name, email, created_at, updated_at
	`, strings.Join(setParts, ", "))

	var user User
	err = db.QueryRow(query, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
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

// query function to fetch user matching email
func (db *DB) GetUserByEmail(email string) (*User, error) {
	query := `
	SELECT id, name, email, created_at, updated_at
	FROM user
	WHERE email = ?
	`

	var user User
	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User with email %s not found", email)
		}
		return nil, fmt.Errorf("Failed to get user: %w", err)
	}

	return &user, nil
}
