package database

import "testing"

func TestDB_ValidateLogin(t *testing.T) {
	// sets up an empty db w/ tables
	db := setupTestDB(t)
	defer teardownTestDB(db)

	// create user first
	newUser := CreateUserRequest{
		Username: "testusername",
		Name:     "testuser",
		Password: "Password123!",
		Email:    "test@example.com",
		Timezone: "America/Vancouver",
	}

	user, _ := db.CreateUser(newUser)

	req := AccountLoginRequest{
		LoginType: "username",
		Username:  newUser.Username,
		Password:  newUser.Password,
	}

	result, err := db.ValidateLogin(req)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	assertField(t, "username", user.Username, result.Username)
}
