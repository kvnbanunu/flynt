package handlers

import (
	"log"
	"net/http"
	"strings"

	"flynt/internal/database"
)

// ProfileHandler handles /profile routes
type ProfileHandler struct {
	db *database.DB
}

func NewProfileHandler(db *database.DB) *ProfileHandler {
	return &ProfileHandler{db: db}
}

func (h *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/profile")

	userIDValue := r.Context().Value("userID")
	userID, ok := userIDValue.(int)

	if !ok {
		writeError(w, http.StatusBadRequest, "Invalid or missing user ID in context")
		return
	}

	switch path {
	case "", "/":
		switch r.Method {
		case http.MethodGet:
			h.getProfile(w, r, userID)
		case http.MethodPost:
			h.updateProfile(w, r, userID)
		default:
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// GET /api/profile
func (h *ProfileHandler) getProfile(w http.ResponseWriter, _ *http.Request, id int) {
	user, err := h.db.GetUserByID(id)
	if err != nil {
		log.Printf("Error getting profile: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get profile")
		return
	}
	writeSuccess(w, http.StatusOK, "Profile retrieved successfully", user)
}

// PUT /api/profile
func (h *ProfileHandler) updateProfile(w http.ResponseWriter, r *http.Request, id int) {
	var req database.UpdateUserRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	user, err := h.db.UpdateUser(id, req)
	if err != nil {
		log.Printf("Error updating profile: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to update profile")
		return
	}

	writeSuccess(w, http.StatusOK, "Profile updated successfully", user)
}

