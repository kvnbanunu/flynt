package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// DB holds the database connection
type DB struct {
	*sqlx.DB
}

// initialize db and return connection
func InitDB(path string) (*DB, error) {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	dbConn := &DB{db}

	if err := dbConn.createTables(); err != nil {
		return nil, fmt.Errorf("Failed to create tables: %w", err)
	}

	log.Println("Database initialized successfully")
	return dbConn, nil
}

// Creates all tables on init
func (db *DB) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		img_url TEXT,
		bio TEXT,
		timezone TEXT DEFAULT 'Canada/Pacific',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TRIGGER IF NOT EXISTS update_user_updated_at
		AFTER UPDATE ON user
	BEGIN
		UPDATE user SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
	END;

	CREATE TABLE IF NOT EXISTS fyre (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		streak_count INTEGER DEFAULT 0,
		user_id INTEGER NOT NULL REFERENCES user(id),
		bonfyre_id INTEGER REFERENCES bonfyre(id),
		active_days TEXT DEFAULT '1111111',
		is_checked INTEGER DEFAULT 0,
		last_checked_at DATETIME,
		last_checked_at_prev DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	
	CREATE TRIGGER IF NOT EXISTS update_fyre_updated_at
		AFTER UPDATE ON fyre
	BEGIN
		UPDATE fyre SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
	END;

	CREATE TABLE IF NOT EXISTS goal_type (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS goal (
		fyre_id  INTEGER NOT NULL REFERENCES fyre(id),
		description TEXT NOT NULL,
		goal_type_id INTEGER NOT NULL REFERENCES goal_type(id),
		data TEXT,
		PRIMARY KEY (fyre_id, description)
	);

	CREATE TABLE IF NOT EXISTS bonfyre (
		id INTEGER PRIMARY KEY AUTOINCREMENT
	);

	CREATE TABLE IF NOT EXISTS friend (
		user_id_1 INTEGER NOT NULL REFERENCES user(id),
		user_id_2 INTEGER NOT NULL REFERENCES user(id),
		status TEXT DEFAULT 'pending',
		PRIMARY KEY (user_id_1, user_id_2)
	);

	CREATE TABLE IF NOT EXISTS social_post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL REFERENCES user(id),
		fyre_id INTEGER NOT NULL REFERENCES fyre(id),
		type TEXT NOT NULL,
		content TEXT NOT NULL
	);
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("Failed to create tables: %w", err)
	}

	return nil
}

// Close db connection
func (db *DB) Close() error {
	return db.DB.Close()
}
