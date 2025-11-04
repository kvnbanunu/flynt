package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type CreateFyreRequest struct {
	Title       string `json:"title"`
	StreakCount int    `json:"streak_count"`
	ActiveDays  string `json:"active_days"`
}

type UpdateFyreRequest struct {
	Title       *string `json:"title"`
	StreakCount *int    `json:"streak_count"`
	BonfyreID   *int    `json:"bonfyre_id"`
	ActiveDays  *string `json:"active_days"`
}

type CheckFyreRequest struct {
	FyreID    int  `json:"id"`
	Increment bool `json:"increment"`
}

func (db *DB) CreateFyre(req CreateFyreRequest, id int) (*Fyre, error) {
	query := `
	INSERT INTO fyre (title, user_id, streak_count, active_days)
	VALUES (?, ?, ?, ?)
	RETURNING *
	`

	var fyre Fyre
	err := db.Get(&fyre, query, req.Title, id, req.StreakCount, req.ActiveDays)
	if err != nil {
		return nil, fmt.Errorf("Failed to create fyre: %w", err)
	}
	return &fyre, nil
}

func (db *DB) GetFyreByID(id int) (*Fyre, error) {
	query := `SELECT * FROM fyre WHERE id = ?`

	var fyre Fyre
	err := db.Get(&fyre, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Fyre not found: %w", err)
		}
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

	if req.Title != nil {
		fields = append(fields, "title = ?")
		args = append(args, req.Title)
	}

	if req.StreakCount != nil {
		fields = append(fields, "streak_count = ?")
		args = append(args, req.StreakCount)
	}

	if req.BonfyreID != nil {
		fields = append(fields, "bonfyre_id = ?")
		args = append(args, req.BonfyreID)
	}

	if req.ActiveDays != nil {
		fields = append(fields, "active_days = ?")
		args = append(args, req.ActiveDays)
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

// resets all fyre checks in list
func (db *DB) ResetChecks(ids []int) ([]Fyre, error) {
	var params []string
	for _, r := range ids {
		params = append(params, fmt.Sprintf("id = %d", r))
	}
	query := fmt.Sprintf(`
	UPDATE fyre
	SET is_checked = false
	WHERE %s
	RETURNING *
	`, strings.Join(params, " OR "))

	var fyres []Fyre
	err := db.Select(&fyres, query)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to reset fyre checks: %w", err)
	}
	return fyres, nil
}

// either increment or decrement streakcount and set last_checked
func (db *DB) CheckFyre(req CheckFyreRequest, id int) (*Fyre, error) {
	query := `
	UPDATE fyre
	SET streak_count = streak_count %s
	is_checked = ?
	WHERE id = ?
	RETURNING *;
	`
	var fyre Fyre
	var err error
	updateLastChecked := `+ 1,
	last_checked_at_prev = last_checked_at,
	last_checked_at = ?,
	`
	isChecked := true
	if req.Increment {
		lastChecked := time.Now()
		query = fmt.Sprintf(query, updateLastChecked)
		err = db.Get(&fyre, query, lastChecked.UTC(), isChecked, req.FyreID)
		if err != nil {
			return nil, err
		}

		// Create the post for daily check
		query = `INSERT into social_post (user_id, fyre_id, type, content)
		VALUES (?, ?, ?, ' just hit a streak of ')
		`
		result, err := db.Exec(query, id, req.FyreID, DailyCheck)
		if err != nil {
			return nil, fmt.Errorf("Failed to create post: %w", err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return nil, fmt.Errorf("Failed to get rows affected: %w", err)
		}

		if rowsAffected == 0 {
			return nil, fmt.Errorf("No post created")
		}
	} else {
		updateLastChecked = `- 1,
		last_checked_at = last_checked_at_prev,
		`
		isChecked = false
		query = fmt.Sprintf(query, updateLastChecked)
		err = db.Get(&fyre, query, isChecked, req.FyreID)
	}
	if err != nil {
		return nil, err
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
