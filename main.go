package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/yaviral17/hw-go/db"
	"github.com/yaviral17/hw-go/myLogs"
)

// LoggingMiddleware logs the details of each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Create a response writer to capture the status code
		rr := &responseRecorder{w, http.StatusOK}
		next.ServeHTTP(rr, r)
		// Determine the color based on the status code
		color := "\033[32m" // Green
		if rr.statusCode >= 400 {
			color = "\033[31m" // Red
		}
		reset := "\033[0m"

		log.Printf(
			"%s[%s] %s %s %s %d %s%s",
			color,
			time.Now().Format(time.RFC3339),
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			rr.statusCode,
			time.Since(start),
			reset,
		)
	})
}

// responseRecorder is a wrapper to capture the status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func main() {

	dbUrl := os.Getenv("DATABASE_URL")

	port := os.Getenv("PORT")

	if dbUrl == "" {
		myLogs.MyErrorLog("DATABASE_URL environment variable is not set")
		return
	}

	// Initialize the database connection
	err := db.InitDB(dbUrl)
	if err != nil {
		myLogs.MyErrorLog("Failed to connect to the database")
		return
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter() // Create a subrouter with the prefix

	apiRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API"))
	})
	// Apply the logging middleware to the router
	loggedRouter := LoggingMiddleware(router)

	http.ListenAndServe("0.0.0.0:"+port, loggedRouter) // Start the server
}
