package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"statusCode"`
	Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Generic error response
func writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	res := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
	json.NewEncoder(w).Encode(res)
}

// Generic success response
func writeSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	w.WriteHeader(statusCode)
	res := SuccessResponse{
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
}

// parse contents of body and store in data
func parseBody(w http.ResponseWriter, r *http.Request, data any) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON payload")
		return err
	}
	return nil
}

// Writes error and returns false if any fields are empty
func validateFields(w http.ResponseWriter, models ...any) bool {
	count := 0
	for _, v := range models {
		switch v := v.(type) {
		case int:
			if v < 0 {
				break
			}
		case string:
			if v == "" {
				break
			}
		default:
			if v == nil {
				break
			}
		}
		count++
	}
	if count < len(models) {
		writeError(w, http.StatusBadRequest, "Missing required fields")
		return false
	}
	return true
}
