package database

import (
	"testing"
)

func TestDB_CreateFyre(t *testing.T) {
	// sets up an empty db w/ tables
	db := setupTestDB(t)
	defer teardownTestDB(db)

	newFyre := CreateFyreRequest{
		Title:       "testfyre",
		StreakCount: 0,
		ActiveDays:  "mon",
		CategoryID:  1,
	}

	fyre, err := db.CreateFyre(newFyre, 1)
	if err != nil {
		t.Fatalf("CreateFyre failed: %v", err)
	}

	assertField(t, "title", newFyre.Title, fyre.Title)
	assertField(t, "category", newFyre.CategoryID, fyre.CategoryID)
	assertField(t, "title", newFyre.ActiveDays, fyre.ActiveDays)
}

func TestDB_UpdateFyre(t *testing.T) {
	// sets up an empty db w/ tables
	db := setupTestDB(t)
	defer teardownTestDB(db)

	// create user first
	newFyre := CreateFyreRequest{
		Title:       "testfyre",
		StreakCount: 0,
		ActiveDays:  "mon",
		CategoryID:  1,
	}

	title := "new title"
	streak := 5
	bon := 3
	days := "mon,tue"
	categoryId := 2
	private := false

	expectedUpdatedFyre := UpdateFyreRequest{
		Title:       &title,
		StreakCount: &streak,
		BonfyreID:   &bon,
		ActiveDays:  &days,
		CategoryID:  &categoryId,
		IsPrivate:   &private,
	}

	_, err := db.CreateFyre(newFyre, 1)
	if err != nil {
		t.Fatalf("CreateFyre failed: %v", err)
	}

	updatedfyre, err := db.UpdateFyre(1, expectedUpdatedFyre)
	if err != nil {
		t.Fatalf("UpdateFyre failed: %v", err)
	}

	updatedfyrefromdb, err := db.GetFyreByID(1)
	if err != nil {
		t.Fatalf("GetFyreByID failed: %v", err)
	}

	assertField(t, "title", updatedfyre.Title, updatedfyrefromdb.Title)
	assertField(t, "category", updatedfyre.CategoryID, updatedfyrefromdb.CategoryID)
	assertField(t, "days", updatedfyre.ActiveDays, updatedfyrefromdb.ActiveDays)
}
