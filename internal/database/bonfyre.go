package database

import (
	"errors"
	"fmt"
)

type BonfyreRequest struct {
	FyreID    int `json:"fyre_id"`
	BonfyreID int `json:"bonfyre_id"`
}

type BonfyreMember struct {
	Username    string `db:"username" json:"username"`
	StreakCount string `db:"streak_count" json:"streak_count"`
}

var ErrAlreadyJoinedBonfyre = errors.New("Already joined BonFyre")

func (db *DB) GetBonfyre(id int) (*Bonfyre, error) {
	query := `SELECT * FROM bonfyre WHERE id = ?`

	var bonfyre Bonfyre
	err := db.Get(&bonfyre, query, id)
	if err != nil {
		return nil, err
	}

	return &bonfyre, nil
}

func (db *DB) GetBonfyreMembers(bonfyreID int) ([]BonfyreMember, error) {
	var members []BonfyreMember

	query := `
	SELECT u.username, f.streak_count FROM fyre f
	JOIN user u ON f.user_id = u.id
	WHERE f.bonfyre_id = ?
	ORDER BY f.streak_count DESC
	`

	err := db.Select(&members, query, bonfyreID)
	if err != nil {
		return members, err
	}

	return members, nil
}

func (db *DB) JoinBonfyre(req BonfyreRequest, userID int) error {
	// get second users fyre
	fyre, err := db.GetFyreByID(req.FyreID)
	if err != nil {
		return err
	}

	var bonfyreID int

	// check if bonfyre exists or create new one
	if fyre.BonfyreID != nil {
		_, err = db.GetBonfyre(*fyre.BonfyreID)
		if err != nil {
			return err
		}

		bonfyreID = *fyre.BonfyreID

		// check if user already joined the bonfyre
		var existing Fyre
		err := db.Get(&existing, `
			SELECT * FROM fyre
			WHERE user_id = ? AND bonfyre_id = ?
			`, userID, bonfyreID)
		if err == nil {
			return ErrAlreadyJoinedBonfyre
		}
	} else {
		var bonfyre Bonfyre
		// new bonfyre based on existing fyre
		err = db.Get(&bonfyre, "INSERT INTO bonfyre (total) VALUES (?) RETURNING *", fyre.StreakCount)
		if err != nil {
			return fmt.Errorf("Failed to create new bonfyre: %w", err)
		}
		bonfyreID = bonfyre.ID

		// add second user to bonfyre
		err = db.Get(fyre, "UPDATE fyre SET bonfyre_id = ? WHERE id = ? RETURNING *", bonfyreID, req.FyreID)
		if err != nil {
			return fmt.Errorf("Failed to update fyre: %w", err)
		}
	}

	// create new fyre for joinee
	err = db.Get(fyre, `
	INSERT INTO fyre (title, streak_count, user_id, bonfyre_id, active_days, category_id)
	VALUES (?, ?, ?, ?, ?, ?)
	RETURNING *
	`, fyre.Title, 0, userID, bonfyreID, fyre.ActiveDays, fyre.CategoryID)
	if err != nil {
		return fmt.Errorf("Failed to add new fyre to bonfyre: %w", err)
	}

	return nil
}
