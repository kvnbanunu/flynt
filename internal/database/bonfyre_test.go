package database

import (
	"testing"
)

func TestDB_JoinBonfyre(t *testing.T) {
	// sets up an empty db w/ tables
	db := setupTestDB(t)
	defer teardownTestDB(db)

	// newUser1 := CreateUserRequest{
	// 	Username: "lemon",
	// 	Name:     "lucas",
	// 	Password: "lemonlemon",
	// 	Email:    "lucas@gmail.com",
	// 	Timezone: "America/Vancouver",
	// }

	// newUser2 := CreateUserRequest{
	// 	Username: "tosen",
	// 	Name:     "evin",
	// 	Password: "tosentosen",
	// 	Email:    "evin@gmail.com",
	// 	Timezone: "America/Vancouver",
	// }

	// _, err := db.CreateUser(newUser1)
	// if err != nil {
	// 	t.Fatalf("CreateUser failed: %v", err)
	// }

	// _, err = db.CreateUser(newUser2)
	// if err != nil {
	// 	t.Fatalf("CreateUser failed: %v", err)
	// }

	// fyre data
	newFyre := CreateFyreRequest{
		Title:       "testfyre",
		StreakCount: 0,
		ActiveDays:  "mon",
		CategoryID:  1,
	}

	// create a fyre in db
	_, err := db.CreateFyre(newFyre, 1)
	if err != nil {
		t.Fatalf("CreateFyre failed: %v", err)
	}

	// create bonfyre data
	newBonFyreRequest := BonfyreRequest{
		FyreID:    1,
		BonfyreID: 1,
	}

	// join
	_ = db.JoinBonfyre(newBonFyreRequest, 1)

	members, err := db.GetBonfyre(1)
	if err != nil {
		t.Fatalf("CreateFyre failed: %v", err)
	}

	// user, _ := db.CreateUser(newUser)

	// req := AccountLoginRequest{
	// 	LoginType: "username",
	// 	Username:  newUser.Username,
	// 	Password:  newUser.Password,
	// }

	// result, err := db.ValidateLogin(req)
	// if err != nil {
	// 	t.Fatalf("Failed to login: %v", err)
	// }

	assertField(t, "id", newBonFyreRequest.BonfyreID, members.ID)
}
