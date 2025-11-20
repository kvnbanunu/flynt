package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"flynt/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

// log http requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)

		log.Printf("%s %s %d %v %s", r.Method, r.URL.Path, rw.statusCode, time.Since(start), r.RemoteAddr)
	})
}

// cors, will need to change once we set prod domain
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", utils.CFG.Client)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := utils.GetConfig()
		cookie, err := r.Cookie(cfg.Context)
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Fprintf(w, "No auth cookie found.\n")
				return
			}
			http.Error(w, "Error retrieving cookie", http.StatusInternalServerError)
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &utils.Claims{}, func(token *jwt.Token) (any, error) {
			return []byte(cfg.Secret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*utils.Claims); ok && token.Valid {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userID", claims.UserID)
			ctx = context.WithValue(ctx, "timezone", claims.Timezone)
			ctx = context.WithValue(ctx, "role", claims.Role)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		}
	})
}

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value("role").(string)
		if role == "admin" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "User is not an admin", http.StatusUnauthorized)
		}
	})
}

// wrap writer to get status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
