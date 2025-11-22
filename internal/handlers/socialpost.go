package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"flynt/internal/database"
)

type SocialPostHandler struct {
	db *database.DB
}

func NewSocialPostHandler(db *database.DB) *SocialPostHandler {
	return &SocialPostHandler{db: db}
}

func (h *SocialPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/socialpost")

	// may change
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userID := r.Context().Value("userID").(int)

	switch {
	case path == "/" || path == "":
		h.GetAllPosts(w, r, userID)
	case strings.HasPrefix(path, "/like"):
		postIDStr := strings.TrimPrefix(path, "/like/")
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			writeError(w, http.StatusNotAcceptable, "Invalid or missing id")
			return
		}
		h.LikePost(w, r, postID)
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

func (h *SocialPostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request, id int) {
	posts, err := h.db.GetAllPosts(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get posts: %v", err))
		return
	}

	writeSuccess(w, http.StatusOK, "Posts retreived successfully", posts)
}

func (h *SocialPostHandler) LikePost(w http.ResponseWriter, r *http.Request, id int) {
	err := h.db.LikePost(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to like post: %v", err))
		return
	}

	writeSuccess(w, http.StatusOK, "Posts liked successfully", nil)
}
