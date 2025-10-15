package handlers

import (
	"log"
	"net/http"
	"strings"

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

	path := strings.TrimPrefix(r.URL.Path, "/account")

	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}

	switch path { // add more later
	case "/register":
		h.register(w, r)
	case "/login":
		h.login(w, r)
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles POST /api/account/register
func (h *AccountHandler) register(w http.ResponseWriter, r *http.Request) {
	var req database.CreateUserRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	if !validateFields(w, req.Username, req.Name, req.Password, req.Email) {
		return
	}

	user, err := h.db.CreateUser(req)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to create") {
			writeError(w, http.StatusNotAcceptable, "Failed to register new user")
			return
		}
		log.Printf("Error register: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to register new user")
		return
	}
	writeSuccess(w, http.StatusOK, "New user registered successfully", user)
}

// Handles POST /api/account/login
func (h *AccountHandler) login(w http.ResponseWriter, r *http.Request) {
	var req database.AccountLoginRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	// maybe validation if necessary here

	user, err := h.db.ValidateLogin(req)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to login") {
			writeError(w, http.StatusNotAcceptable, "Email or Password incorrect")
			return
		}
		log.Printf("Error validating login: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to login")
		return
	}
	writeSuccess(w, http.StatusOK, "User logged in successfully", user)
}
