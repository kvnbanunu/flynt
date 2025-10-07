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

	err = db.insertDummyFriends()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) insertDummyUsers() error {
	dummyUser := os.Getenv("DUMMY_USER")
	dummypass := os.Getenv("DUMMY_PASSWORD")
	adminID := os.Getenv("ADMIN_ID")
	adminUser := os.Getenv("ADMIN_USER")
	adminPass := os.Getenv("ADMIN_PASSWORD")

	adminPassHashed, err := utils.HashPassword(adminPass)
	if err != nil {
		return err
	}

	hashed, err := utils.HashPassword(dummypass)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
	INSERT INTO user (name, email, password)
	VALUES
	('Test User', '%s', '%s'),
	('Brandon Rada', 'brandon@gmail.com', '%s'),
	('Evin Gonzales', 'evin@gmail.com', '%s'),
	('Lucas Laviolette', 'lucas@gmail.com', '%s'),
	('Kevin Nguyen', 'kevin@gmail.com', '%s')
	`, dummyUser, hashed, hashed, hashed, hashed, hashed)

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert dummy users: %w", err)
	}

	query = fmt.Sprintf(`
	INSERT INTO user (id, name, email, password)
	VALUES
	(%s, 'Admin User', '%s', '%s')
	`, adminID, adminUser, adminPassHashed)

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert admin user: %w", err)
	}

	return nil
}

func (db *DB) insertDummyFyres() error {
	query := `
	INSERT INTO fyre (title, streak_count, user_id, active_days)
	VALUES
	('Win a game of Clash Royale', 100, 2, '10000000'),
	('Drink water', 0, 3, '10000000'),
	('Gacha', 99, 3, '10000000'),
	('Open Pokemon TCG pack', 9, 4, '00000001'),
	('Go for a run', 2, 5, '00101010');
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert dummy fyres: %w", err)
	}

	return nil
}

func (db *DB) insertDummyFriends() error {
	query := `
	INSERT INTO friend (user_id_1, user_id_2, status)
	VALUES
	(2, 3, 'friends'),
	(2, 5, 'pending'),
	(3, 2, 'friends'),
	(3, 4, 'friends'),
	(4, 3, 'friends'),
	(4, 5, 'friends'),
	(5, 3, 'blocked'),
	(5, 2, 'sent'),
	(5, 4, 'friends')
	`
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to insert dummy friends: %w", err)
	}

	return nil
}
