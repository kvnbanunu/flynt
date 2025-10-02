package handlers

import (
	"encoding/json"
	"flynt/internal/database"
	"log"
	"net/http"
	"strings"
)

type FyreHandler struct {
	db *database.DB
}

func NewFyreHandler(db *database.DB) *FyreHandler {
	return &FyreHandler{db: db}
}

func (h *FyreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (h *FyreHandler) createFyre(w http.ResponseWriter, r *http.Request) {
	var req database.CreateFyreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid Json payload")
		return
	}

	if req.Title == "" || req.UserID == 0 {
		h.writeError(w, http.StatusBadRequest, "Title and user_id are required")
		return
	}

	fyre, err := h.db.CreateFyre(req)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			h.writeError(w, http.StatusBadRequest, "Fyre with this title already exists")
			return
		}
		log.Printf("Error creating fyre: %v", err)
		h.writeError(w, http.StatusBadRequest, "Failed to create fyre")
		return
	}
	h.writeSuccess(w, http.StatusCreated, "Fyre created successfully", fyre)
}

func (h *FyreHandler) getFyreById(w http.ResponseWriter, r *http.Request) {

}

func (h *FyreHandler) getAllUserFyres(w http.ResponseWriter, r *http.Request) {

}

func (h *FyreHandler) updateFyre(w http.ResponseWriter, r *http.Request) {

}

func (h *FyreHandler) deleteFyre(w http.ResponseWriter, r *http.Request) {

}

// error response for FyreHandlers
func (h *FyreHandler) writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	res := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
	json.NewEncoder(w).Encode(res)
}

// success response for FyreHandlers
func (h *FyreHandler) writeSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	w.WriteHeader(statusCode)
	res := SuccessResponse{
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
}
