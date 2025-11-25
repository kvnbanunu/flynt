package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"flynt/internal/database"
	"flynt/internal/utils"
)

type FyreHandler struct {
	db *database.DB
}

func NewFyreHandler(db *database.DB) *FyreHandler {
	return &FyreHandler{db: db}
}

func (h *FyreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/fyre")

	if path == "/categories" || path == "/categories/" {
		if r.Method == http.MethodGet {
			h.getAllCategories(w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
		return
	}

	userID := r.Context().Value("userID").(int)
	timezone := r.Context().Value("timezone").(string)

	switch {
	case path == "" || path == "/":
		switch r.Method {
		case http.MethodGet:
			h.getAllUserFyres(w, r, userID, timezone)
		case http.MethodPost:
			h.createFyre(w, r, userID)
		default:
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	case strings.HasPrefix(path, "/"):
		switch {
		case path == "/check" || path == "/check/":
			if r.Method != http.MethodPut {
				writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
				return
			}
			h.checkFyre(w, r, userID)
		case strings.HasPrefix(path, "/bonfyre"):
			if strings.HasPrefix(path, "/bonfyre/") && r.Method == http.MethodGet {
				bonfyreIDStr := strings.TrimPrefix(path, "/bonfyre/")
				bonfyreID, err := strconv.Atoi(bonfyreIDStr)
				if err != nil {
					writeError(w, http.StatusBadRequest, "Bonfyre id missing or invalid", err)
					return
				}
				h.getBonfyreMembers(w, r, bonfyreID)
				return
			}
			switch r.Method {
			case http.MethodGet:
				h.getBonfyre(w, r)
			case http.MethodPost:
				h.joinBonfyre(w, r, userID)
			default:
				writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
			}
		case path == "/full" || path == "full/":
			if r.Method != http.MethodGet {
				writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
				return
			}
			h.getAllUserFullFyres(w, r, userID)
		case strings.HasPrefix(path, "/user/"):
			idStr := strings.TrimPrefix(path, "/user/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				writeError(w, http.StatusBadRequest, "Invalid user ID")
				return
			}
			if r.Method != http.MethodGet {
				writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
				return
			}
			h.getAllFriendsFyres(w, r, id)
		default:
			idStr := strings.TrimPrefix(path, "/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				writeError(w, http.StatusBadRequest, "Invalid fyre ID")
				return
			}
			switch r.Method {
			case http.MethodGet:
				h.getFyreById(w, r, id)
			case http.MethodPut:
				h.updateFyre(w, r, id)
			case http.MethodDelete:
				h.deleteFyre(w, r, id)
			default:
				writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
			}
		}
	default:
		writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

func (h *FyreHandler) getAllCategories(w http.ResponseWriter, _ *http.Request) {
	categories, err := h.db.GetAllCategories()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get categories")
		return
	}
	writeSuccess(w, http.StatusOK, "Categories retrieved successfully", categories)
}

// handles POST /fyre
func (h *FyreHandler) createFyre(w http.ResponseWriter, r *http.Request, id int) {
	var req database.CreateFyreRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	if !utils.ValidateFields(req.Title) {
		writeError(w, http.StatusBadRequest, "Title is required")
		return
	}

	fyre, err := h.db.CreateFyre(req, id)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			writeError(w, http.StatusConflict, "Fyre with this title already exists")
			return
		}
		log.Printf("Error creating fyre: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to create fyre")
		return
	}
	writeSuccess(w, http.StatusCreated, "Fyre created successfully", fyre)
}

// handles GET /fyre/{id}
func (h *FyreHandler) getFyreById(w http.ResponseWriter, _ *http.Request, id int) {
	fyre, err := h.db.GetFyreByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error getting fyre: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get fyre")
		return
	}
	writeSuccess(w, http.StatusOK, "Fyre retreived successfully", fyre)
}

// handles PUT /api/fyre/{id}
func (h *FyreHandler) updateFyre(w http.ResponseWriter, r *http.Request, id int) {
	var req database.UpdateFyreRequest

	if err := parseBody(w, r, &req); err != nil {
		return
	}

	fyre, err := h.db.UpdateFyre(id, req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error updating fyre: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to update fyre")
		return
	}
	writeSuccess(w, http.StatusOK, "Fyre updated successfully", fyre)
}

// hanldes DELETE /fyre/{id}
func (h *FyreHandler) deleteFyre(w http.ResponseWriter, _ *http.Request, id int) {
	err := h.db.DeleteFyre(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to delete Fyre", err)
		return
	}
	writeSuccess(w, http.StatusOK, "Fyre successfully deleted", nil)
}

// Handles PUT /fyre/check
func (h *FyreHandler) checkFyre(w http.ResponseWriter, r *http.Request, id int) {
	var req database.CheckFyreRequest
	err := parseBody(w, r, &req)
	if err != nil {
		return
	}

	existingFyre, err := h.db.GetFyreByID(req.FyreID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "Fyre not found")
			return
		}
		log.Printf("Error fetching fyre: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to fetch fyre")
		return
	}

	// if unchecking we need to set the last_checked date to yesterday
	if req.Increment && existingFyre.IsChecked ||
		!req.Increment && !existingFyre.IsChecked {
		writeError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	checkedFyre, err := h.db.CheckFyre(req, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to check Fyre")
		return
	}

	existingFyre.IsChecked = checkedFyre.IsChecked
	existingFyre.StreakCount = checkedFyre.StreakCount
	existingFyre.LastCheckedAt = checkedFyre.LastCheckedAt
	existingFyre.LastCheckedAtPrev = checkedFyre.LastCheckedAtPrev

	writeSuccess(w, http.StatusOK, "Fyre updated successfully", existingFyre)
}

func (h *FyreHandler) getAllUserFullFyres(w http.ResponseWriter, _ *http.Request, id int) {
	fyres, err := h.db.GetAllUserFyres(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get fyres", err)
		return
	}

	fullFyres, err := h.db.MapFyreGoals(fyres, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to reset fyre checks", err)
		return
	}
	writeSuccess(w, http.StatusOK, "Fyres retrieved successfully", fullFyres)
}

// handles GET /fyre
// Finds all fyres that belong to this user id
// You only want to do this once per day / login ***
func (h *FyreHandler) getAllUserFyres(w http.ResponseWriter, _ *http.Request, id int, timezone string) {
	fyres, err := h.db.GetAllUserFyres(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get fyres", err)
		return
	}

	var toUpdate []int
	var indexes []int
	for i, f := range fyres {
		if f.LastCheckedAt == nil || !f.IsChecked {
			continue
		}
		passed, err := utils.CheckDayPassed(*f.LastCheckedAt, timezone)
		if err != nil {
			writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to check fyre dates %v", err))
			return
		}
		if passed {
			toUpdate = append(toUpdate, f.ID)
			indexes = append(indexes, i)
		}
	}

	if len(toUpdate) > 0 {
		updated, err := h.db.ResetChecks(toUpdate)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to reset fyre checks")
			return
		}
		// remap fyres list
		for i := range updated {
			fyres[indexes[i]].IsChecked = updated[i].IsChecked
		}
	}

	fullFyres, err := h.db.MapFyreGoals(fyres, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to reset fyre checks", err)
		return
	}
	writeSuccess(w, http.StatusOK, "Fyres retrieved successfully", fullFyres)
}

// Get /fyre/user/{id} shallow get no updates
func (h *FyreHandler) getAllFriendsFyres(w http.ResponseWriter, _ *http.Request, id int) {
	fyres, err := h.db.GetAllFriendsFyres(id)
	if err != nil {
		log.Printf("Error getting fyres: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to get fyres")
		return
	}

	writeSuccess(w, http.StatusOK, "Fyres retrieved successfully", fyres)
}
