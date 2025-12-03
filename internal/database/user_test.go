package database

import (
	"fmt"
	"testing"
)

func TestDB_CreateUser(t *testing.T) {
	// sets up an empty db w/ tables
	db := setupTestDB(t)
	defer teardownTestDB(db)

	newUser := CreateUserRequest{
		Username: "lemon",
		Name:     "lucas",
		Password: "lemonlemon",
		Email:    "lucas@gmail.com",
		Timezone: "America/Vancouver",
	}

	result, err := db.CreateUser(newUser)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	assertField(t, "username", newUser.Username, result.Username)
	assertField(t, "name", newUser.Name, result.Name)
	assertField(t, "password", "", result.Password) // security reasons
	assertField(t, "title", newUser.Email, result.Email)
	assertField(t, "title", newUser.Timezone, result.Timezone)
}

func TestDB_GetUserByID(t *testing.T) {
	// sets up an empty db w/ tables
	db := setupTestDB(t)
	defer teardownTestDB(db)

	newUser := CreateUserRequest{
		Username: "lemon",
		Name:     "lucas",
		Password: "lemonlemon",
		Email:    "lucas@gmail.com",
		Timezone: "America/Vancouver",
	}

	result, err := db.CreateUser(newUser)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	userResult, err := db.GetUserByID(1)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	fmt.Println(userResult)
	assertField(t, "username", userResult.Username, result.Username)
	assertField(t, "name", userResult.Name, result.Name)
	assertField(t, "title", userResult.Email, result.Email)
	assertField(t, "title", userResult.Timezone, result.Timezone)
}

func TestDB_UpdateUser(t *testing.T) {
	// sets up an empty db w/ tables
	db := setupTestDB(t)
	defer teardownTestDB(db)

	newUser := CreateUserRequest{
		Username: "lemon",
		Name:     "lucas",
		Password: "lemonlemon",
		Email:    "lucas@gmail.com",
		Timezone: "America/Vancouver",
	}

	name := "evin"
	currentPassword := "lemonlemon"
	newPassword := "tosentosen"
	email := "evin@gmail.com"
	imgURL := ""
	bio := "live laugh love"
	timezone := "America/Toronto"

	newUpdatedUser := UpdateUserRequest{
		Name:            &name,
		CurrentPassword: &currentPassword,
		NewPassword:     &newPassword,
		Email:           &email,
		ImgURL:          &imgURL,
		Bio:             &bio,
		Timezone:        &timezone,
	}

	_, err := db.CreateUser(newUser)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	userUpdated, err := db.UpdateUser(1, newUpdatedUser)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	fmt.Println(userUpdated)
	assertField(t, "name", userUpdated.Name, name)
	assertField(t, "email", userUpdated.Email, email)
	assertField(t, "timezone", userUpdated.Timezone, timezone)
}
