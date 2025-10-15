package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"flynt/internal/database"
)

type FriendRequest struct {
	Type string `json:"type"`
	ID1  int    `json:"user_id_1"`
	ID2  int    `json:"user_id_2"`
}

type FriendHandler struct {
	db *database.DB
}

func NewFriendHandler(db *database.DB) *FriendHandler {
	return &FriendHandler{db: db}
}

func (h *FriendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/friend")

	switch {
	case path == "" || path == "/":
		switch r.Method {
		case http.MethodPut:
			h.updateFriend(w, r)
		case http.MethodPost:
			h.addFriend(w, r)
		case http.MethodDelete:
			h.deleteFriend(w, r)
		default:
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/"):
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		idStr := strings.TrimPrefix(path, "/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}
		h.getFriendsList(w, r, id)
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles POST /api/friend
func (h *FriendHandler) addFriend(w http.ResponseWriter, r *http.Request) {
	var req FriendRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	if req.Type != "addfriend" {
		writeError(w, http.StatusBadRequest, "Incorrect method and request type combination")
		return
	}

	update := database.UpdateFriendRequest{ID1: req.ID1, ID2: req.ID2}

	err := h.db.AddFriend(update)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to add friend")
		return
	}
	writeSuccess(w, http.StatusCreated, "Friend request created", nil)
}

// handles PUT /api/friend
func (h *FriendHandler) updateFriend(w http.ResponseWriter, r *http.Request) {
	var req FriendRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	update := database.UpdateFriendRequest{ID1: req.ID1, ID2: req.ID2}

	switch req.Type {
	case "acceptfriend":
		err := h.db.AcceptFriend(update)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to accept friend request")
			return
		}
		writeSuccess(w, http.StatusOK, "Friend request accepted", nil)
	case "blockfriend":
		err := h.db.BlockFriend(update)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to block friend")
			return
		}
		writeSuccess(w, http.StatusOK, "Friend successfully blocked", nil)
	default:
		writeError(w, http.StatusBadRequest, "Incorrect method and request type combination")
	}
}

// handles DELETE /api/friend
func (h *FriendHandler) deleteFriend(w http.ResponseWriter, r *http.Request) {
	var req FriendRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}
	if req.Type != "deletefriend" {
		writeError(w, http.StatusBadRequest, "Incorrect method and request type combination")
		return
	}

	update := database.UpdateFriendRequest{ID1: req.ID1, ID2: req.ID2}

	err := h.db.DeleteFriend(update)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to remove friend")
		return
	}
	writeSuccess(w, http.StatusOK, "Friend successfully removed", nil)
}

// handles GET /api/friend/{user_id_1}
func (h *FriendHandler) getFriendsList(w http.ResponseWriter, _ *http.Request, id int) {
	friendsList, err := h.db.GetFriendsList(id)
	if err != nil {
		log.Printf("Error getting friendslist: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get friendslist")
		return
	}
	writeSuccess(w, http.StatusOK, "FriendsList retrieved successfully", friendsList)
}
