package handlers

// This handler is used to perform health checks during runtime

import (
	"encoding/json"
	"net/http"
	"time"

	"flynt/internal/database"
)

type HealthHandler struct {
	db *database.DB
}

type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Checks    map[string]string `json:"checks"`
}

func NewHealthHandler(db *database.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "Method Not Allowed",
			Message: "Only GET method is allowed for health checks",
		})
		return
	}

	checks := make(map[string]string)
	status := "healthy"

	// check db status
	if err := h.db.Ping(); err != nil {
		checks["database"] = "unhealthy: " + err.Error()
		status = "unhealthy"
	} else {
		checks["database"] = "healthy"
	}

	// -----> future checks here <-----

	checks["server"] = "healthy"

	response := HealthResponse{
		Status:    status,
		Timestamp: time.Now(),
		Checks:    checks,
	}

	w.Header().Set("Content-Type", "application/json")
	if status == "unhealthy" {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
}
