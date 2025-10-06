package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"flynt/internal/database"
	"flynt/internal/handlers"
	"flynt/internal/utils"
)

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// init handlers
	userHandler := handlers.NewUserHandler(db)
	accountHandler := handlers.NewAccountHandler(db)
	fyreHandler := handlers.NewFyreHandler(db)
	friendHandler := handlers.NewFriendHandler(db)
	healthHandler := handlers.NewHealthHandler(db)

	// init server router
	mux := http.NewServeMux()

	// routes
	mux.Handle("/api/user", userHandler)
	mux.Handle("/api/user/", userHandler)
	mux.Handle("/api/account/", accountHandler)
	mux.Handle("/api/fyre", fyreHandler)
	mux.Handle("/api/fyre/", fyreHandler)
	mux.Handle("/api/fyre/user/", fyreHandler)
	mux.Handle("/api/friend", friendHandler)
	mux.Handle("/api/friend/", friendHandler)
	mux.Handle("/health", healthHandler)

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

	// setup server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      loggingMiddleware(corsMiddleware(mux)),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// start server in goroutine
	go func() {
		log.Printf("Server starting on port%s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// wait for Signal Interrupt for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("\nServer is shutting down...")

	// timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
		return
	}

	log.Println("Server exited successfully")
}

// Middleware, may move elsewhere

// log http requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)

		log.Printf("%s %s %d %v %s", r.Method, r.URL.Path, rw.statusCode, time.Since(start), r.RemoteAddr)
	})
}

// cors, will need to change once we set prod domain
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("ENV") == "production" {
			w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_URL"))
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
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
