package database

import (
	"time"
)

// Represents user table in db
type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Name      string    `db:"name" json:"name"`
	Password  string    `db:"password" json:"password,omitempty"`
	Email     string    `db:"email" json:"email,omitempty"`
	ImgURL    *string   `db:"img_url" json:"img_url,omitempty"`
	Bio       *string   `db:"bio" json:"bio,omitempty"`
	Timezone  string    `db:"timezone" json:"timezone,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Represents fyre table in db
type Fyre struct {
	ID                int        `db:"id" json:"id"`
	Title             string     `db:"title" json:"title"`
	StreakCount       int        `db:"streak_count" json:"streak_count"`
	UserID            int        `db:"user_id" json:"user_id"`
	BonfyreID         *int       `db:"bonfyre_id" json:"bonfyre_id"`
	ActiveDays        string     `db:"active_days" json:"active_days"`
	IsChecked         bool       `db:"is_checked" json:"is_checked"`
	LastCheckedAt     *time.Time `db:"last_checked_at" json:"last_checked_at,omitempty"`
	LastCheckedAtPrev *time.Time `db:"last_checked_at_prev" json:"last_checked_at_prev,omitempty"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updated_at"`
	CategoryID	  int        `db:"category_id" json:"category_id"`
}

// Represents category table in db
type Category struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// Represents goal_type table in db
type GoalType struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// Represents goal table in db
type Goal struct {
	FyreID      int     `db:"fyre_id" json:"fyre_id"`
	Description string  `db:"description" json:"description"`
	GoalTypeID  int     `db:"goal_type_id" json:"goal_type_id"`
	Data        *string `db:"data" json:"data"`
}

// Represents bonfyre table in db
type Bonfyre struct {
	ID int `db:"id" json:"id"`
}

// Represents friend table in db
type Friend struct {
	UserID1 int    `db:"user_id_1" json:"user_id_1,omitempty"`
	UserID2 int    `db:"user_id_2" json:"user_id_2"`
	Status  string `db:"status" json:"status"`
}

type FriendStatus string

const (
	Pending FriendStatus = "pending"
	Sent    FriendStatus = "sent"
	Friends FriendStatus = "friends"
	Blocked FriendStatus = "blocked"
)
