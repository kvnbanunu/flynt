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

	db.AddFriend(newFriendRequest)
	
	res, err := db.GetFriendsList(r1.ID)
	assertField(t, "username", r2.Username, res[0].Username)

}
