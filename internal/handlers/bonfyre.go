package handlers

import (
	"errors"
	"net/http"

	"flynt/internal/database"
)

func (h *FyreHandler) getBonfyre(w http.ResponseWriter, r *http.Request) {
}

// POST /fyre/bonfyre
func (h *FyreHandler) joinBonfyre(w http.ResponseWriter, r *http.Request, id int) {
	var req database.BonfyreRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	err := h.db.JoinBonfyre(req, id)
	if err != nil {
		if errors.Is(err, database.ErrAlreadyJoinedBonfyre) {
			writeError(w, http.StatusBadRequest, "Already joined BonFyre", err)
			return
		}
		writeError(w, http.StatusBadRequest, "Failed to join bonfyre", err)
		return
	}

	writeSuccess(w, http.StatusOK, "Bonfyre joined successfully", nil)
}

func (h *FyreHandler) getBonfyreMembers(w http.ResponseWriter, _ *http.Request, bonfyreID int) {
	members, err := h.db.GetBonfyreMembers(bonfyreID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get bonfyre members", err)
		return
	}

	writeSuccess(w, http.StatusOK, "Bonfyre members retrieved successfully", members)
}
