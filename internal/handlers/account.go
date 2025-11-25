package handlers

import (
	"log"
	"net/http"
	"strings"

	"flynt/internal/database"
	"flynt/internal/utils"
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
	case "/register", "/register/":
		h.register(w, r)
	case "/login", "/login/":
		h.login(w, r)
	case "/logout", "/logout/":
		h.logout(w, r)
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles POST /account/register
func (h *AccountHandler) register(w http.ResponseWriter, r *http.Request) {
	var req database.CreateUserRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	if !utils.ValidateFields(req.Username, req.Name, req.Password, req.Email, req.Timezone) {
		writeError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	user, err := h.db.CreateUser(req)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to create") {
			writeError(w, http.StatusNotAcceptable, "Failed to register new user")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to register new user", err)
		return
	}

	err = setCookie(w, user.ID, user.Timezone)
	if err != nil {
		return
	}

	writeSuccess(w, http.StatusOK, "New user registered successfully", user)
}

// Handles POST /account/login
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
		writeError(w, http.StatusInternalServerError, "Failed to login", err)
		return
	}

	err = setCookie(w, user.ID, user.Timezone)
	if err != nil {
		return
	}

	writeSuccess(w, http.StatusOK, "User logged in successfully", user)
}

// Handles POST /account/logout (just deletes auth cookie)
func (h *AccountHandler) logout(w http.ResponseWriter, _ *http.Request) {
	setExpiredCookie(w)
	writeSuccess(w, http.StatusOK, "User logged out successfully", nil)
}
