package handlers

import (
	"log"
	"net/http"
	"strings"

	"flynt/internal/database"
)

type FriendRequest struct {
	Type string `json:"type"`
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

	path := strings.TrimPrefix(r.URL.Path, "/friend")
	id := r.Context().Value("userID").(int)

	switch path {
	case "", "/":
		switch r.Method {
		case http.MethodPut:
			h.updateFriend(w, r, id)
		case http.MethodPost:
			h.addFriend(w, r, id)
		case http.MethodDelete:
			h.deleteFriend(w, r, id)
		default:
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case "/all", "/all/":
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		
		h.getFriendsList(w, r, id)
	case "/non", "/non/":
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		h.getNonFriendsList(w, r, id)
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// handles POST /friend
func (h *FriendHandler) addFriend(w http.ResponseWriter, r *http.Request, id int) {
	var req FriendRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	if req.Type != "addfriend" {
		writeError(w, http.StatusBadRequest, "Incorrect method and request type combination")
		return
	}

	update := database.UpdateFriendRequest{ID1: id, ID2: req.ID2}

	err := h.db.AddFriend(update)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to add friend")
		return
	}
	writeSuccess(w, http.StatusCreated, "Friend request created", nil)
}

// handles PUT /friend
func (h *FriendHandler) updateFriend(w http.ResponseWriter, r *http.Request, id int) {
	var req FriendRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	update := database.UpdateFriendRequest{ID1: id, ID2: req.ID2}

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

// handles DELETE /friend
func (h *FriendHandler) deleteFriend(w http.ResponseWriter, r *http.Request, id int) {
	var req FriendRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}
	if req.Type != "deletefriend" {
		writeError(w, http.StatusBadRequest, "Incorrect method and request type combination")
		return
	}

	update := database.UpdateFriendRequest{ID1: id, ID2: req.ID2}

	err := h.db.DeleteFriend(update)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to remove friend")
		return
	}
	writeSuccess(w, http.StatusOK, "Friend successfully removed", nil)
}

// handles GET /friend/all
func (h *FriendHandler) getFriendsList(w http.ResponseWriter, _ *http.Request, id int) {
	friendsList, err := h.db.GetFriendsList(id)
	if err != nil {
		log.Printf("Error getting friendslist: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get friendslist")
		return
	}
	writeSuccess(w, http.StatusOK, "FriendsList retrieved successfully", friendsList)
}

// handles GET /friend/non
func (h *FriendHandler) getNonFriendsList(w http.ResponseWriter, _ *http.Request, id int) {
	friendsList, err := h.db.GetNonFriendsList(id)
	if err != nil {
		log.Printf("Error getting friendslist: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get friendslist")
		return
	}
	writeSuccess(w, http.StatusOK, "FriendsList retrieved successfully", friendsList)
}
