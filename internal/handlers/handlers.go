package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"flynt/internal/database"
	"flynt/internal/middleware"
	"flynt/internal/utils"
)

type ErrorResponse struct {
	Error   string `json:"status_code"`
	Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SetupHandlers(db *database.DB) http.Handler {
	mux := http.NewServeMux()

	// init handlers
	userHandler := NewUserHandler(db)
	accountHandler := NewAccountHandler(db)
	fyreHandler := NewFyreHandler(db)
	friendHandler := NewFriendHandler(db)
	healthHandler := NewHealthHandler(db)
	goalHandler := NewGoalHandler(db)
	socialPostHandler := NewSocialPostHandler(db)

	// alias for readability
	auth := middleware.Auth
	admin := middleware.Admin

	// routes
	mux.Handle("/user", auth(userHandler))
	mux.Handle("/user/", auth(userHandler))
	mux.Handle("/account/", accountHandler)
	mux.Handle("/goal", auth(goalHandler))
	mux.Handle("/goal/", auth(goalHandler))
	mux.Handle("/fyre", auth(fyreHandler))
	mux.Handle("/fyre/", auth(fyreHandler))
	mux.Handle("/friend", auth(friendHandler))
	mux.Handle("/friend/", auth(friendHandler))
	mux.Handle("/health", auth(admin(healthHandler)))
	mux.Handle("/socialpost", auth(socialPostHandler))
	mux.Handle("/socialpost/", auth(socialPostHandler))

	// root handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// placeholder, will change later
		msg := `{"message":"API Server is running","version":"1.0.0","endpoints":["/health","/api/users"]}`
		w.Write([]byte(msg))
	})

	return middleware.Logger(middleware.Cors(mux))
}

func setCookie(w http.ResponseWriter, id int, timezone string) error {
	cfg := utils.GetConfig()
	token, err := utils.NewToken(id, timezone)
	if err != nil {
		fmt.Println("Failed to create token")
		writeError(w, http.StatusInternalServerError, "Failed to create token")
		return err
	}

	isProduction := cfg.Env == "production"

	cookie := &http.Cookie{
		Name:     cfg.Context,
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		Path:     "/",
		Secure:   isProduction,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	if isProduction {
		cookie.Domain = cfg.Domain
	}

	http.SetCookie(w, cookie)
	return nil
}

func setExpiredCookie(w http.ResponseWriter) {
	cfg := utils.GetConfig()

	isProduction := cfg.Env == "production"

	cookie := &http.Cookie{
		Name:     cfg.Context,
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Secure:   isProduction,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	if isProduction {
		cookie.Domain = cfg.Domain
	}

	http.SetCookie(w, cookie)
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
