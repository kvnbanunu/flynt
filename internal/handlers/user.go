package handlers

import (
	"log"
	"net/http"
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

	id := r.Context().Value("userID").(int)

	path := strings.TrimPrefix(r.URL.Path, "/user")

	switch {
	case path == "" || path == "/":
		switch r.Method {
		case http.MethodGet:
			h.getUserByID(w, r, id)
		case http.MethodPost: // use account/register instead
			h.createUser(w, r)
		case http.MethodPut:
			h.updateUser(w, r, id)
		case http.MethodDelete:
			h.deleteUser(w, r, id)
		default:
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/"):
		path = strings.TrimPrefix(path, "/")
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		switch path {
		case "all", "all/":
			h.getAllUsers(w, r, 0, false)
		case "redacted", "redacted/":
			h.getAllUsers(w, r, id, true)
		default:
			writeError(w, http.StatusNotFound, "Endpoint not found")
		}
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles	POST /user
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

// handles	GET /user/all or /user/redacted
func (h *UserHandler) getAllUsers(w http.ResponseWriter, _ *http.Request, id int, redacted bool) {
	var users []database.User
	var err error
	if redacted {
		users, err = h.db.GetAllUsersRedacted(id)
	} else {
		users, err = h.db.GetAllUsers()
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get users", err)
		return
	}

	writeSuccess(w, http.StatusOK, "Users retrieved successfully", users)
}

// handles	GET /user/
func (h *UserHandler) getUserByID(w http.ResponseWriter, _ *http.Request, id int) {
	user, err := h.db.GetUserByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "User not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to get user", err)
		return
	}

	user.Password = ""

	writeSuccess(w, http.StatusOK, "User retrieved successfully", user)
}

// handles	PUT /user
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
		if strings.Contains(err.Error(), "Incorrect password") || strings.Contains(err.Error(), "Both password") {
			writeError(w, http.StatusBadRequest, err.Error(), err)
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	user.Password = ""

	writeSuccess(w, http.StatusOK, "User updated successfully", user)
}

// handles	DELETE /users
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
