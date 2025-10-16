package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"flynt/internal/database"
	"flynt/internal/utils"
)

// handles requests for user ops
type UserHandler struct {
	db *database.DB
}

func NewUserHandler(db *database.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/user")

	switch {
	case path == "" || path == "/":
		switch r.Method {
		case http.MethodGet:
			h.getAllUsers(w, r)
		case http.MethodPost:
			h.createUser(w, r)
		default:
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/"): // id queries
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.getUserByID(w, r, id)
		case http.MethodPut:
			h.updateUser(w, r, id)
		case http.MethodDelete:
			h.deleteUser(w, r, id)
		default:
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles	POST /api/user
func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req database.CreateUserRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	// validation
	if !utils.ValidateFields(req.Username, req.Name, req.Password, req.Email) {
		writeError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	user, err := h.db.CreateUser(req)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			writeError(w, http.StatusConflict, "User with this name or email already exists")
			return
		}
		log.Printf("Error creating user: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	writeSuccess(w, http.StatusCreated, "User created successfully", user)
}

// handles	GET /api/user
func (h *UserHandler) getAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := h.db.GetAllUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	writeSuccess(w, http.StatusOK, "Users retrieved successfully", users)
}

// handles	GET /api/user/{id}
func (h *UserHandler) getUserByID(w http.ResponseWriter, _ *http.Request, id int) {
	user, err := h.db.GetUserByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error getting user: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get user")
		return
	}

	writeSuccess(w, http.StatusOK, "User retrieved successfully", user)
}

// handles	PUT /api/user/{id}
func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var req database.UpdateUserRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	user, err := h.db.UpdateUser(id, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error updating user: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	writeSuccess(w, http.StatusOK, "User updated successfully", user)
}

// handles	DELETE /api/users/{id}
func (h *UserHandler) deleteUser(w http.ResponseWriter, _ *http.Request, id int) {
	err := h.db.DeleteUser(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error deleting user: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	writeSuccess(w, http.StatusOK, "User successfully deleted", nil)
}
