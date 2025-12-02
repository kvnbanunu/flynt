package handlers

import (
	"net/http"
	"testing"

	"flynt/internal/database"
)

func TestAccount_Register_Success(t *testing.T) {
	// create a mock database
	mockDB := database.NewMockDB()

	// define expected behaviour
	user := &database.User{
		ID:        1,
		Username:  "testusername",
		Name:      "testuser",
		Password:  "Password123!",
		Email:     "test@example.com",
		ImgURL:    nil,
		Bio:       nil,
		Timezone:  "America/Vancouver",
		FyreTotal: 0,
	}

	// mock db request - returns expected
	mockDB.CreateUserFunc = func(req database.CreateUserRequest) (*database.User, error) {
		return user, nil
	}

	// the mock handler to test
	handler := NewAccountHandler(mockDB)

	// Create request
	reqBody := database.CreateUserRequest{
		Username: "testusername",
		Name:     "testuser",
		Password: "Password123!",
		Email:    "test@example.com",
		Timezone: "America/Vancouver",
	}

	res := mockRequest(t, http.MethodPost, "/account/register", http.StatusOK, reqBody, handler.register)

	// validate relevant fields
	assertField(t, "id", user.ID, res)
	assertField(t, "username", user.Username, res)
}
