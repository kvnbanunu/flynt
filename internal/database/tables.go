package database

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Image     string    `json:"img_url"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Fyre struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	UserID      int       `json:"user_id"`
	StreakCount int       `json:"streak_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Goal struct {
	FyreID int `json:"fyre_id"`
	Name   int `json:"name"`
}

type Bonfyre struct {
	ID      int `json:"id"`
	FyreID1 int `json:"fyre_id_1"`
	FyreID2 int `json:"fyre_id_2"`
}

type Friend struct {
	UserID1 int `json:"user_id_1"`
	UserID2 int `json:"user_id_2"`
}
