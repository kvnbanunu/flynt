package handlers

import (
	"net/http"
	"testing"

	"flynt/internal/database"
)

func TestGet_User_By_ID(t *testing.T) {
	mockDB := database.NewMockDB()

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

	mockDB.GetUserByIDFunc = func(id int) (*database.User, error) {
		return user, nil
	}

	handler := NewUserHandler(mockDB)

	res := mockRequestWithID(t, http.MethodGet, "/user/1", http.StatusOK, nil, handler.getUserByID, user.ID)

	assertField(t, "id", user.ID, res)
	assertField(t, "username", user.Username, res)
}
