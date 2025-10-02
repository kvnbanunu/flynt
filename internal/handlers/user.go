package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"flynt/internal/database"
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

	path := strings.TrimPrefix(r.URL.Path, "/api/user")

	switch {
	case path == "" || path == "/":
		switch r.Method {
		case http.MethodGet:
			h.getAllUsers(w, r)
		case http.MethodPost:
			h.createUser(w, r)
		default:
			h.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/"): // id queries
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.writeError(w, http.StatusBadRequest, "Invalid user ID")
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
			h.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	default:
		h.writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles	POST /api/user
func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req database.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// validation
	if req.Name == "" || req.Email == "" || req.Password == "" {
		h.writeError(w, http.StatusBadRequest, "Name, email, and password are required")
		return
	}

	user, err := h.db.CreateUser(req)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			h.writeError(w, http.StatusConflict, "User with this email already exists")
			return
		}
		log.Printf("Error creating user: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	h.writeSuccess(w, http.StatusCreated, "User created successfully", user)
}

// handles	GET /api/user
func (h *UserHandler) getAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := h.db.GetAllUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	h.writeSuccess(w, http.StatusOK, "Users retrieved successfully", users)
}

// handles	GET /api/user/{id}
func (h *UserHandler) getUserByID(w http.ResponseWriter, _ *http.Request, id int) {
	user, err := h.db.GetUserByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.writeError(w, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error getting user: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to get user")
		return
	}

	h.writeSuccess(w, http.StatusOK, "User retrieved successfully", user)
}

// handles	PUT /api/user/{id}
func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var req database.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	user, err := h.db.UpdateUser(id, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.writeError(w, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error updating user: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	h.writeSuccess(w, http.StatusOK, "User updated successfully", user)
}

// handles	DELETE /api/users/{id}
func (h *UserHandler) deleteUser(w http.ResponseWriter, _ *http.Request, id int) {
	err := h.db.DeleteUser(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.writeError(w, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error deleting user: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}
}

// error response for UserHandlers
func (h *UserHandler) writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	res := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
	json.NewEncoder(w).Encode(res)
}

// success response for UserHandlers
func (h *UserHandler) writeSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	w.WriteHeader(statusCode)
	res := SuccessResponse{
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
}
