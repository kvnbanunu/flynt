package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"flynt/internal/database"
)

type GoalHandler struct {
	db *database.DB
}

func NewGoalHandler(db *database.DB) *GoalHandler {
	return &GoalHandler{db: db}
}

func (h *GoalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/goal")
	switch {
	case path == "", path == "/":
		if r.Method == http.MethodPost {
			h.createGoal(w, r)
			return
		}
	case strings.HasPrefix(path, "/"):
		idStr := strings.TrimPrefix(path, "/")
		fyreID, err := strconv.Atoi(idStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid fyre ID")
			return
		}
		switch r.Method {
		case http.MethodGet:
			h.getGoalByFyre(w, r, fyreID)
		case http.MethodPut:
			h.updateGoal(w, r, fyreID)
		case http.MethodDelete:
			h.deleteGoal(w, r, fyreID)
		default:
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles POST /api/goal
func (h *GoalHandler) createGoal(w http.ResponseWriter, r *http.Request) {
	var req database.CreateGoalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if req.FyreID == 0 || req.GoalTypeID == 0 {
		writeError(w, http.StatusBadRequest, "fyre_id and goal_type_id are required")
		return
	}

	goal, err := h.db.CreateGoal(req)
	if err != nil {
		log.Printf("Error creating goal: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to create goal")
		return
	}

	writeSuccess(w, http.StatusCreated, "Goal created successfully", goal)
}

// GET /api/goal/fyre/{fyre_id}
func (h *GoalHandler) getGoalByFyre(w http.ResponseWriter, _ *http.Request, fyreID int) {
	goal, err := h.db.GetGoalByID(fyreID)
	if err != nil {
		log.Printf("Error getting goal: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get goal")
		return
	}

	if goal == nil {
		writeError(w, http.StatusNotFound, "Goal not found for this fyre")
		return
	}

	writeSuccess(w, http.StatusOK, "Goal retrieved successfully", goal)
}

// PUT /api/goal/{fyre_id}
func (h *GoalHandler) updateGoal(w http.ResponseWriter, r *http.Request, fyreID int) {
	var req database.UpdateGoalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	goal, err := h.db.UpdateGoal(fyreID, req)
	if err != nil {
		log.Printf("Error updating goal: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to update goal")
		return
	}

	writeSuccess(w, http.StatusOK, "Goal updated successfully", goal)
}

func (h *GoalHandler) deleteGoal(w http.ResponseWriter, _ *http.Request, fyreID int) {
	err := h.db.DeleteGoal(fyreID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "Goal not found for this fyre")
			return
		}
		log.Printf("Error deleting goal: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to delete goal")
		return
	}

	writeSuccess(w, http.StatusOK, "Goal deleted successfully", nil)
}
