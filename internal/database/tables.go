package database

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImgURL    string    `json:"img_url"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Fyre struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	UserID      int       `json:"user_id"`
	StreakCount int       `json:"streak_count"`
	BonfyreID   int       `json:"bonfyre_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Goal struct {
	FyreID      int    `json:"fyre_id"`
	Description string `json:"description"`
	Name        int    `json:"name"`
}

type Bonfyre struct {
	ID      int `json:"id"`
}

type Friend struct {
	UserID1 int `json:"user_id_1"`
	UserID2 int `json:"user_id_2"`
}
