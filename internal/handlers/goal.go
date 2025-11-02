package handlers

import (
	"encoding/json"
	"fmt"
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

	path := r.URL.Path
	fmt.Printf("REQUEST PATH: %q\n", path)

	if strings.HasPrefix(path, "/goal/") {
		idStr := strings.TrimPrefix(path, "/goal/")
		fyreID, err := strconv.Atoi(idStr)
		if err != nil {
			h.writeError(w, http.StatusBadRequest, "Invalid fyre ID")
			return
		}

		// Handle methods
		switch r.Method {
		case http.MethodGet:
			h.getGoalByFyre(w, r, fyreID)
		case http.MethodPost:
			h.createGoal(w, r)
		case http.MethodPut:
			h.updateGoal(w, r, fyreID)
		case http.MethodDelete:
			h.deleteGoal(w, r, fyreID)
		default:
			h.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
		return
	}

	// Path not found
	h.writeError(w, http.StatusNotFound, "Endpoint not found")
}

// handles POST /api/goal
func (h *GoalHandler) createGoal(w http.ResponseWriter, r *http.Request) {
	var req database.CreateGoalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	if req.FyreID == 0 || req.GoalTypeID == 0 {
		h.writeError(w, http.StatusBadRequest, "fyre_id and goal_type_id are required")
		return
	}

	goal, err := h.db.CreateGoal(req)
	if err != nil {
		log.Printf("Error creating goal: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to create goal")
		return
	}

	h.writeSuccess(w, http.StatusCreated, "Goal created successfully", goal)
}

// GET /api/goal/fyre/{fyre_id}
func (h *GoalHandler) getGoalByFyre(w http.ResponseWriter, _ *http.Request, fyreID int) {
	goal, err := h.db.GetGoalByID(fyreID)
	if err != nil {
		log.Printf("Error getting goal: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to get goal")
		return
	}

	if goal == nil {
		h.writeError(w, http.StatusNotFound, "Goal not found for this fyre")
		return
	}

	h.writeSuccess(w, http.StatusOK, "Goal retrieved successfully", goal)
}

// PUT /api/goal/{fyre_id}
func (h *GoalHandler) updateGoal(w http.ResponseWriter, r *http.Request, fyreID int) {
	var req database.UpdateGoalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	goal, err := h.db.UpdateGoal(fyreID, req)
	if err != nil {
		log.Printf("Error updating goal: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to update goal")
		return
	}

	h.writeSuccess(w, http.StatusOK, "Goal updated successfully", goal)
}

func (h *GoalHandler) deleteGoal(w http.ResponseWriter, _ *http.Request, fyreID int) {
	err := h.db.DeleteGoal(fyreID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.writeError(w, http.StatusNotFound, "Goal not found for this fyre")
			return
		}
		log.Printf("Error deleting goal: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to delete goal")
		return
	}

	h.writeSuccess(w, http.StatusOK, "Goal deleted successfully", nil)
}

// error response for GoalHandlers
func (h *GoalHandler) writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	res := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
	json.NewEncoder(w).Encode(res)
}

// success response for GoalHandlers
func (h *GoalHandler) writeSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	w.WriteHeader(statusCode)
	res := SuccessResponse{
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
}
