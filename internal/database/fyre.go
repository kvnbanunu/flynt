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
	CategoryID  int    `json:"category_id"`
}

type UpdateFyreRequest struct {
	Title       *string `json:"title"`
	StreakCount *int    `json:"streak_count"`
	BonfyreID   *int    `json:"bonfyre_id"`
	ActiveDays  *string `json:"active_days"`
	CategoryID  *int    `json:"category_id"`
	IsPrivate   *bool   `json:"is_private"`
}

type CheckFyreRequest struct {
	FyreID    int  `json:"id"`
	Increment bool `json:"increment"`
}

type FyreTotalResponse struct {
	FyreTotal int `db:"fyre_total" json:"fyre_total"`
}

// used for left joins when no goals are found
type GoalJoin struct {
	FyreID      int     `db:"fyre_id" json:"fyre_id"`
	Description *string `db:"description" json:"description"`
	GoalTypeID  *int    `db:"goal_type_id" json:"goal_type_id"`
	Data        *string `db:"data" json:"data"`
}

type FullFyre struct {
	Fyre  Fyre       `json:"fyre"`
	Goals []GoalJoin `json:"goals,omitempty"`
}

func (db *DB) GetAllCategories() ([]Category, error) {
	query := `SELECT * FROM category ORDER BY name ASC`
	var categories []Category
	err := db.Select(&categories, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	return categories, nil
}

func (db *DB) CreateFyre(req CreateFyreRequest, id int) (*Fyre, error) {
	query := `
	INSERT INTO fyre (title, user_id, streak_count, active_days, category_id)
	VALUES (?, ?, ?, ?, ?)
	RETURNING *
	`

	var fyre Fyre
	err := db.Get(&fyre, query, req.Title, id, req.StreakCount, req.ActiveDays, req.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("Failed to create fyre: %w", err)
	}

	if req.StreakCount > 0 {
		_, err = db.UpdateFyreTotal(id, req.StreakCount, true)
		if err != nil {
			return nil, fmt.Errorf("Failed to add to fyre total")
		}
	}

	return &fyre, nil
}

func (db *DB) GetFyreByID(id int) (*Fyre, error) {
	query := `
	SELECT fyre.*, bonfyre.total AS bonfyre_total
	FROM fyre
	LEFT JOIN bonfyre ON fyre.bonfyre_id = bonfyre.id
	WHERE fyre.id = ?`

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

// Gets and maps goals to the full fyre array
func (db *DB) MapFyreGoals(fyres []Fyre, userID int) ([]FullFyre, error) {
	var fullFyres []FullFyre
	var goals []GoalJoin

	query := `
	SELECT fyre.id AS fyre_id, goal.description, goal.goal_type_id, goal.data
	FROM fyre
	LEFT JOIN goal
	ON fyre.id = goal.fyre_id
	WHERE fyre.user_id = ?
	ORDER BY fyre.id DESC
	`

	err := db.Select(&goals, query, userID)
	if err != nil {
		return fullFyres, err
	}

	var collapsed [][]GoalJoin
	for i, v := range goals {
		// check if this goal is connected to prev fyre
		if i == 0 || goals[i-1].FyreID != v.FyreID {
			collapsed = append(collapsed, []GoalJoin{v})
			continue
		}
		collapsed[i-1] = append(collapsed[i-1], v)
	}

	// map each goal array to proper fyre
	for i := range fyres {
		fullFyres = append(fullFyres, FullFyre{Fyre: fyres[i]})
		if collapsed[i][0].Description != nil {
			fullFyres[i].Goals = collapsed[i]
		}
	}
	return fullFyres, nil
}

// Get all fyres for just this user, returns empty array if none
func (db *DB) GetAllUserFyres(id int) ([]Fyre, error) {
	query := `
	SELECT fyre.*, bonfyre.total AS bonfyre_total
	FROM fyre
	LEFT JOIN bonfyre ON fyre.bonfyre_id = bonfyre.id
	WHERE user_id = ?
	ORDER BY fyre.id DESC`

	var fyres []Fyre
	err := db.Select(&fyres, query, id)
	if err != nil {
		return fyres, fmt.Errorf("Failed to get fyres: %w", err)
	}
	return fyres, nil
}

// only get relevant fields
func (db *DB) GetAllFriendsFyres(id int) ([]Fyre, error) {
	query := `SELECT id, title, streak_count, bonfyre_id, is_checked, category_id
	FROM fyre
	WHERE user_id = ? AND is_private = false
	ORDER BY streak_count DESC
	`

	var fyres []Fyre
	err := db.Select(&fyres, query, id)
	if err != nil {
		return fyres, fmt.Errorf("Failed to get fyres: %w", err)
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

	if req.CategoryID != nil {
		fields = append(fields, "category_id = ?")
		args = append(args, req.CategoryID)
	}

	if req.IsPrivate != nil {
		fields = append(fields, "is_private = ?")
		args = append(args, req.IsPrivate)
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

		if !fyre.IsPrivate {
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

	_, err = db.UpdateFyreTotal(fyre.UserID, 1, req.Increment)
	if err != nil {
		return nil, err
	}

	if fyre.BonfyreID != nil {
		err = db.incrementBonfyreTotal(*fyre.BonfyreID, req.Increment)
		if err != nil {
			return nil, err
		}
	}

	return &fyre, nil
}

func (db *DB) incrementBonfyreTotal(bonfyreID int, increment bool) error {
	var bonfyre Bonfyre
	update := "+ 1"
	if !increment {
		update = "- 1"
	}
	query := fmt.Sprintf(`
		UPDATE bonfyre SET
		total = total %s
		WHERE id = ?
		RETURNING *
		`, update)
	err := db.Get(&bonfyre, query, bonfyreID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetFyreTotal(id int) (*FyreTotalResponse, error) {
	query := `SELECT fyre_total FROM user WHERE id = ?`

	var total FyreTotalResponse
	err := db.Get(&total, query, id)
	if err != nil {
		return nil, err
	}

	return &total, err
}

func (db *DB) UpdateFyreTotal(id, amount int, increment bool) (*FyreTotalResponse, error) {
	operator := "+"
	if !increment {
		operator = "-"
	}
	query := fmt.Sprintf(`
	UPDATE user
	SET fyre_total = fyre_total %s ?
	WHERE id = ?
	RETURNING fyre_total
	`, operator)

	var total FyreTotalResponse
	err := db.Get(&total, query, amount, id)
	if err != nil {
		return nil, err
	}

	return &total, err
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
