package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type CreateGoalRequest struct {
	FyreID      int     `json:"fyre_id"`
	Description string  `json:"description"`
	GoalTypeID  int     `json:"goal_type_id"`
	Data        *string `json:"data"`
}

type UpdateGoalRequest struct {
	Description string  `json:"description"`
	GoalTypeID  int     `json:"goal_type_id"`
	Data        *string `json:"data"`
}

func (db *DB) CreateGoal(req CreateGoalRequest) (*Goal, error) {
	query := `
	INSERT INTO goal (fyre_id, description, goal_type_id, data)
	VALUES (?, ?, ?, ?)
	RETURNING *
	`

	var goal Goal
	err := db.Get(&goal, query, req.FyreID, req.Description, req.GoalTypeID, req.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to create goal: %w", err)
	}

	return &goal, nil
}

func (db *DB) GetGoalByID(fyreID int) (*Goal, error) {
	query := `SELECT * FROM goal WHERE fyre_id = ?`

	var goal Goal
	err := db.Get(&goal, query, fyreID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to get goal: %w", err)
	}

	return &goal, nil
}

func (db *DB) UpdateGoal(fyreID int, req UpdateGoalRequest) (*Goal, error) {
	existingGoal, err := db.GetGoalByID(fyreID)
	if err != nil {
		return nil, err
	}
	if existingGoal == nil {
		return nil, fmt.Errorf("No goal found for fyre_id %d", fyreID)
	}

	fields := []string{}
	args := []any{}

	if req.Description != "" {
		fields = append(fields, "description = ?")
		args = append(args, req.Description)
	}
	if req.GoalTypeID != 0 {
		fields = append(fields, "goal_type_id = ?")
		args = append(args, req.GoalTypeID)
	}
	if req.Data != nil {
		fields = append(fields, "data = ?")
		args = append(args, req.Data)
	}

	if len(fields) == 0 {
		return existingGoal, nil
	}

	args = append(args, fyreID)
	query := fmt.Sprintf(`
		UPDATE goal
		SET %s
		WHERE fyre_id = ?
		RETURNING *
	`, strings.Join(fields, ", "))

	var goal Goal
	err = db.Get(&goal, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to update goal: %w", err)
	}

	return &goal, nil
}

func (db *DB) DeleteGoal(fyreID int) error {
	query := `DELETE FROM goal WHERE fyre_id = ?`
	result, err := db.Exec(query, fyreID)
	if err != nil {
		return fmt.Errorf("Failed to delete goal: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("No goal deleted (fyre_id %d not found)", fyreID)
	}

	return nil
}
