package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"flynt/internal/database"
)

type FyreHandler struct {
	db *database.DB
}

func NewFyreHandler(db *database.DB) *FyreHandler {
	return &FyreHandler{db: db}
}

func (h *FyreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/fyre")

	switch {
	case path == "" || path == "/":
		if r.Method == http.MethodPost {
			h.createFyre(w, r)
		} else {
			h.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/user/"):
		if r.Method == http.MethodGet {
			idStr := strings.TrimPrefix(path, "/user/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				h.writeError(w, http.StatusBadRequest, "Invalid user ID")
				return
			}
			h.getAllUserFyres(w, r, id)
		} else {
			h.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/"):
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.writeError(w, http.StatusBadRequest, "Invalid fyre ID")
			return
		}
		switch r.Method {
		case http.MethodGet:
			h.getFyreById(w, r, id)
		case http.MethodPut:
			h.updateFyre(w, r, id)
		case http.MethodDelete:
			h.deleteFyre(w, r, id)
		default:
			h.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	default:
		h.writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles POST /api/fyre
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
		if strings.Contains(err.Error(), "already exists") {
			h.writeError(w, http.StatusConflict, "Fyre with this title already exists")
			return
		}
		log.Printf("Error creating fyre: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to create fyre")
		return
	}
	h.writeSuccess(w, http.StatusCreated, "Fyre created successfully", fyre)
}

// handles GET /api/fyre/{id}
func (h *FyreHandler) getFyreById(w http.ResponseWriter, _ *http.Request, id int) {
	fyre, err := h.db.GetFyreByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error getting fyre: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to get fyre")
		return
	}
	h.writeSuccess(w, http.StatusOK, "Fyre retreived successfully", fyre)
}

// handles PUT /api/fyre/{id}
func (h *FyreHandler) updateFyre(w http.ResponseWriter, r *http.Request, id int) {
	var req database.UpdateFyreRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	fyre, err := h.db.UpdateFyre(id, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error updating fyre: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to update fyre")
		return
	}
	h.writeSuccess(w, http.StatusOK, "Fyre updated successfully", fyre)
}

// hanldes DELETE /api/fyre/{id}
func (h *FyreHandler) deleteFyre(w http.ResponseWriter, _ *http.Request, id int) {
	err := h.db.DeleteFyre(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error deleting fyre: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to delete Fyre")
		return
	}
}

// handles GET /api/fyre/user/{id}
// Finds all fyres that belong to this user id
func (h *FyreHandler) getAllUserFyres(w http.ResponseWriter, _ *http.Request, id int) {
	fyres, err := h.db.GetAllUserFyres(id)
	if err != nil {
		log.Printf("Error getting fyres: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to get fyres")
		return
	}
	h.writeSuccess(w, http.StatusOK, "Fyres retrieved successfully", fyres)
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
