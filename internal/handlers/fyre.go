package handlers

import (
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
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/user/"):
		if r.Method == http.MethodGet {
			idStr := strings.TrimPrefix(path, "/user/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				writeError(w, http.StatusBadRequest, "Invalid user ID")
				return
			}
			h.getAllUserFyres(w, r, id)
		} else {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/"):
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid fyre ID")
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
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles POST /api/fyre
func (h *FyreHandler) createFyre(w http.ResponseWriter, r *http.Request) {
	var req database.CreateFyreRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	if req.Title == "" || req.UserID == 0 {
		writeError(w, http.StatusBadRequest, "Title and user_id are required")
		return
	}

	fyre, err := h.db.CreateFyre(req)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			writeError(w, http.StatusConflict, "Fyre with this title already exists")
			return
		}
		log.Printf("Error creating fyre: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to create fyre")
		return
	}
	writeSuccess(w, http.StatusCreated, "Fyre created successfully", fyre)
}

// handles GET /api/fyre/{id}
func (h *FyreHandler) getFyreById(w http.ResponseWriter, _ *http.Request, id int) {
	fyre, err := h.db.GetFyreByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error getting fyre: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get fyre")
		return
	}
	writeSuccess(w, http.StatusOK, "Fyre retreived successfully", fyre)
}

// handles PUT /api/fyre/{id}
func (h *FyreHandler) updateFyre(w http.ResponseWriter, r *http.Request, id int) {
	var req database.UpdateFyreRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	fyre, err := h.db.UpdateFyre(id, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error updating fyre: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to update fyre")
		return
	}
	writeSuccess(w, http.StatusOK, "Fyre updated successfully", fyre)
}

// hanldes DELETE /api/fyre/{id}
func (h *FyreHandler) deleteFyre(w http.ResponseWriter, _ *http.Request, id int) {
	err := h.db.DeleteFyre(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error deleting fyre: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to delete Fyre")
		return
	}
	writeSuccess(w, http.StatusOK, "Fyre successfully deleted", nil)
}

// handles GET /api/fyre/user/{id}
// Finds all fyres that belong to this user id
func (h *FyreHandler) getAllUserFyres(w http.ResponseWriter, _ *http.Request, id int) {
	fyres, err := h.db.GetAllUserFyres(id)
	if err != nil {
		log.Printf("Error getting fyres: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get fyres")
		return
	}
	writeSuccess(w, http.StatusOK, "Fyres retrieved successfully", fyres)
}
