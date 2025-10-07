package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"flynt/internal/database"
)

type GoalHandler struct (
	db *database.DB
)

func NewGoalHandler(db *database.DB) *GoalHandler {
	return &GoalHandler{db: db} 
}

func (h *GoalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/goal")

	switch {
	case path == "" || path == "/":
		if r.Method == http.MethodPost {
			h.createGoal(w,r)
		}
	}
}
