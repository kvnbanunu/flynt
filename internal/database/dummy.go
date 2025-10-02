package database

import (
	"fmt"
	"os"
	"flynt/internal/utils"
)

func (db *DB) InsertDummyData() error {
	err := db.insertDummyUsers()
	if err != nil {
		return err
	}

	err = db.insertDummyFyres()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) insertDummyUsers() error {
	dummypass := os.Getenv("DUMMY_PASSWORD")

	hashed, err := utils.HashPassword(dummypass)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
	INSERT INTO user (name, email, password)
	VALUES
	('Brandon Rada', 'brandon@gmail.com', '%s'),
	('Evin Gonzales', 'evin@gmail.com', '%s'),
	('Lucas Laviolette', 'lucas@gmail.com', '%s'),
	('Kevin', 'kevin@gmail.com', '%s');
	`, hashed, hashed, hashed, hashed)

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert dummy users: %w", err)
	}

	return nil
}

func (db *DB) insertDummyFyres() error {
	query := `
	INSERT INTO fyre (title, streak_count, user_id, active_days)
	VALUES
	('Win a game of Clash Royale', 100, 1, '10000000'),
	('Drink water', 0, 2, '10000000'),
	('Gacha', 99, 2, '10000000'),
	('Open Pokemon TCG pack', 9, 3, '00000001'),
	('Go for a run', 2, 4, '00101010');
	`

	if _,err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert dummy fyres: %w", err)
	}

	return nil
}
