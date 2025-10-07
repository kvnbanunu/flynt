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

	path := strings.TrimPrefix(r.URL.Path, "/api/goal")

	switch {
	case path == "" || path == "/":
		if r.Method == http.MethodPost {
			h.createGoal(w,r)
		} else {
			h.writeError( w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/fyre/"):
		if r.method == http.MethodGet {
			idStr := strings.TrimPrefix(path, "/fyre/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				h.writeError(w, http.StatusBadRequest, "Invalid Fyre ID")
				return
			}
			h.getAllFyreGoals(w, r, id)
		} else {
			h.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/"):
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.writeError(w, http.StatusBadRequest, "Invalid goal ID")
		}
		switch r.Method {
		case http.MethodGet:
			h.getGoalsByFyre(w, r, id)
		case http.MethodPut:
			h.updateGoal(w, r, id)
		case http.MethodDelete:
			h.deleteGoal(w, r, id)
		default: 
			h.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	default:
		h.writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles POST /api/goal
func (h *GoalHandler) createGoal(w http.ResponseWriter, r *http.Request) {
	var req database.CreateGoalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid Json Payload")
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

func (h *GoalHandler) getAllFyreGoals(w http.ResponseWriter, _ *http.Request, id int) {
	goals, err := h.db.GetGoalsByFyre(id)
	if err != nil {
		if strings.Contains(err.Error(), "fyre not found") {
			h.writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		if strings.Contains(err.Error(), "goals not found") {
			h.writeError(w, http.StatusNotFound, "Goals not found")
			return
		}
		log.Printf("Error getting fyre: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failedto get fyre")
		return
	}
	h.writeSuccess(w, http.StatusOK, "Goals retreived successfully", goals)
}

func (h *GoalHandler) updateGoal(w http.ResponseWriter, r *http.Request, fyre_id int, type_id int) {

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
