package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB holds the database connection
type DB struct {
	*sql.DB
}

// initialize db and return connection
func InitDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite3", path)
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
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		img_url TEXT,
		bio TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TRIGGER IF NOT EXISTS update_users_updated_at
		AFTER UPDATE ON user
	BEGIN
		UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
	END;
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
