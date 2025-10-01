package database

import (
	"fmt"
	"os"
	"flynt/internal/utils"
)

func (db *DB) InsertDummyData() error {
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
	('Lucas Laviolette', 'lucas@gmai.com', '%s'),
	('Kevin', 'kevin@gmail.com', '%s');
	`, hashed, hashed, hashed, hashed)

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert dummy data: %w", err)
	}

	return nil
}
