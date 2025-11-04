package handlers

import (
	"fmt"
	"net/http"
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

	switch path {
	case "/", "":
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		h.GetAllPosts(w, r)
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

func (h *SocialPostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.db.GetAllPosts()
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get posts: %v", err))
		return
	}

	writeSuccess(w, http.StatusOK, "Posts retreived successfully", posts)
}
