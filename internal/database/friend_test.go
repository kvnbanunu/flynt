package database

import (
	"testing"
)

func TestDB_AddFriend(t *testing.T) {
	// sets up an empty db w/ tables
	db := setupTestDB(t)
	defer teardownTestDB(db)

	newUser1 := CreateUserRequest{
		Username: "lemon",
		Name:     "lucas",
		Password: "lemonlemon",
		Email:    "lucas@gmail.com",
		Timezone: "America/Vancouver",
	}

	newUser2 := CreateUserRequest{
		Username: "tosen",
		Name:     "evin",
		Password: "tosentosen",
		Email:    "evin@gmail.com",
		Timezone: "America/Vancouver",
	}

	r1, err := db.CreateUser(newUser1)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	r2, err := db.CreateUser(newUser2)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	newFriendRequest := UpdateFriendRequest{
		ID1: r1.ID,
		ID2: r2.ID,
	}

	acceptFriendRequest := UpdateFriendRequest{
		ID1: r2.ID,
		ID2: r1.ID,
	}

	db.AddFriend(newFriendRequest)

	// assertField(t, "username", newUser.Username, result.Username)
	// assertField(t, "name", newUser.Name, result.Name)
	// assertField(t, "password", "", result.Password) // security reasons
	// assertField(t, "title", newUser.Email, result.Email)
	// assertField(t, "title", newUser.Timezone, result.Timezone)
}
