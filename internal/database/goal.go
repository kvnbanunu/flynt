package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type CreateGoalRequest struct {
	FyreID int `json:"fyre_id"`
	Description string `json:"description"`
	GoalTypeID int `json:"goal_type_id"`
	Data *string `json:"data"`
}

type UpdateFyreRequest struct {
	FyreID int `json:"fyre_id"`
	Description string `json:"description"`
	GoalTypeID int `json:"goal_type_id"`
	Data *string `json:"data"`
}

func (db *DB) CreateFyre(req CreateFyreRequest) (*Fyre, error) {

}
