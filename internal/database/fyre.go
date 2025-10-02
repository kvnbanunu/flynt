package database

import (
	"fmt"
	"strings"
	"time"
)

type CreateFyreRequest struct {
	Title       string `json:"title"`
	StreakCount int    `json:"streak_count,omitempty"`
	UserID      int    `json:"user_id"`
	ActiveDays  string `json:"active_days,omitempty"`
}

type UpdateFyreRequest struct{
	Title       string `json:"title,omitempty"`
	StreakCount int    `json:"streak_count"` // set to -1 if no change
	BonfyreID   int    `json:"bonfyre_id"` // set to -1 if no change
	ActiveDays  string `json:"active_days,omitempty"`
}

func (db *DB) CreateFyre(req CreateFyreRequest) (*Fyre, error) {
	fields := []string{"title", "user_id"}
	placeholders := []string{"?", "?"}
	args := []any{req.Title, req.UserID}

	if req.StreakCount != 0 {
		fields = append(fields, "streak_count")
		placeholders = append(placeholders, "?")
		args = append(args, req.StreakCount)
	}
	
	if req.ActiveDays != "" {
		fields = append(fields, "active_days")
		placeholders = append(placeholders, "?")
		args = append(args, req.ActiveDays)
	}

	query := fmt.Sprintf(`
	INSERT INTO fyre (%s)
	VALUES (%s)
	RETURNING *
	`, strings.Join(fields, ", "), strings.Join(placeholders, ", "))

	var fyre Fyre
	err := db.Get(&fyre, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to create fyre: %w", err)
	}
	return &fyre, nil
}

func (db *DB) GetFyreByID(id int) (*Fyre, error) {
	query := `SELECT * FROM fyre WHERE id = ?`

	var fyre Fyre
	err := db.Select(&fyre, query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get fyre: %w", err)
	}
	return &fyre, nil
}

// Get all fyres for just this user, returns empty array if none
func (db *DB) GetAllUserFyres(id int) ([]Fyre, error) {
	query := `SELECT * FROM fyre WHERE user_id = ? ORDER BY created_at DESC`

	var fyres []Fyre
	err := db.Select(&fyres, query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get fyres: %w", err)
	}
	return fyres, nil
}

func (db *DB) UpdateFyre(id int, req UpdateFyreRequest) (*Fyre, error) {
	existingFyre, err := db.GetFyreByID(id)
	if err != nil {
		return nil, err
	}

	fields := []string{}
	args := []any{}

	if req.Title != "" {
		fields = append(fields, "title = ?")
		args = append(args, req.Title)
	}


	if req.StreakCount > -1 {
		fields = append(fields, "title = ?")
		args = append(args, req.Title)
	}


	if req.BonfyreID > -1 {
		fields = append(fields, "title = ?")
		args = append(args, req.Title)
	}


	if req.ActiveDays != "" {
		fields = append(fields, "title = ?")
		args = append(args, req.Title)
	}

	if len(fields) == 0 { // no change
		return existingFyre, nil
	}

	fields = append(fields, "updated_at = ?")
	args = append(args, time.Now())
	args = append(args, id)
	
	query := fmt.Sprintf(`
		UPDATE fyre
		SET %s
		WHERE id = ?
		RETURNING *
	`, strings.Join(fields, ","))

	var fyre Fyre
	err = db.Get(&fyre, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to update fyre: %w", err)
	}

	return &fyre, nil
}

func (db *DB) DeleteFyre(id int) error {
	_, err := db.GetFyreByID(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM fyre WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete fyre: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No fyre deleted")
	}

	return nil
}
