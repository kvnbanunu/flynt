package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"log"

	"flynt/internal/database"
)

type AccountHandler struct {
	db *database.DB
}

func NewAccountHandler(db *database.DB) *AccountHandler {
	return &AccountHandler{db: db}
}

func (h *AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/account")

	switch  path { // add more later
	case "/login":
		h.login(w, r)
	default:
		h.writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// Handles POST /api/account/login
func (h *AccountHandler) login(w http.ResponseWriter, r *http.Request) {
	var req database.AccountLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// maybe validation if necessary here

	user, err := h.db.ValidateLogin(req)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to login") {
			h.writeError(w, http.StatusNotAcceptable, "Email or Password incorrect")
			return
		}
		log.Printf("Error validating login: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to login")
		return
	}
	h.writeSuccess(w, http.StatusOK, "User logged in successfully", user)
}

// error response for UserHandlers
func (h *AccountHandler) writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	res := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
	json.NewEncoder(w).Encode(res)
}

// success response for UserHandlers
func (h *AccountHandler) writeSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	w.WriteHeader(statusCode)
	res := SuccessResponse{
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
}
