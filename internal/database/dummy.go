package database

import "fmt"

func (db *DB) InsertDummyData() error {
	query := `
	INSERT INTO user (name, email)
	VALUES
	('Brandon Rada', 'brandon@gmail.com'),
	('Evin Gonzales', 'evin@gmail.com'),
	('Lucas Laviolette', 'lucas@gmai.com'),
	('Kevin', 'kevin@gmail.com');
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert dummy data: %w", err)
	}

	return nil
}
