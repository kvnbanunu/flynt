package database

import "time"

// Represents user table in db
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	ImgURL    string    `json:"img_url"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Represents fyre table in db
type Fyre struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	StreakCount int       `json:"streak_count"`
	UserID      int       `json:"user_id"`
	BonfyreID   int       `json:"bonfyre_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Represents goal_type table in db
type GoalType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Represents goal table in db
type Goal struct {
	FyreID      int    `json:"fyre_id"`
	Description string `json:"description"`
	GoalTypeID  int    `json:"goal_type_id"`
	Data        string `json:"data"`
}

// Represents bonfyre table in db
type Bonfyre struct {
	ID int `json:"id"`
}

// Represents friend table in db
type Friend struct {
	UserID1 int `json:"user_id_1"`
	UserID2 int `json:"user_id_2"`
}
