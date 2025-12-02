package database

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// creates in memory db for testing
func setupTestDB(t *testing.T) *DB {
	database, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	db := &DB{database}

	if err := db.createTables(); err != nil {
		t.Fatalf("Failed to create tables: %v", err)
	}
	return db
}

func teardownTestDB(db *DB) {
	if db.DB != nil {
		db.DB.Close()
	}
}

func assertField(t *testing.T, name string, expected, result any) {
	if expected != result {
		t.Errorf("Expected %s %v, got %v", name, expected, result)
	}
}
