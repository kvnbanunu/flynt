package handlers

import (
	"flynt/internal/database"
	"fmt"
	"net/http"
)

func (h *FyreHandler) getBonfyre(w http.ResponseWriter, r *http.Request) {

}

// POST /fyre/bonfyre
func (h *FyreHandler) joinBonfyre(w http.ResponseWriter, r *http.Request, id int) {
	var req database.JoinBonfyreRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	err := h.db.JoinBonfyre(req, id)
	if err != nil {
		fmt.Printf("Error: %v", err)
		writeError(w, http.StatusBadRequest, "Failed to join bonfyre")
		return
	}

	writeSuccess(w, http.StatusOK, "Bonfyre joined successfully", nil)
}
