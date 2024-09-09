package myLogs

import (
	"log"
	"net/http"
	"time"
)

func MySuccessLog(message string) {
	green := "\033[32m"
	reset := "\033[0m"
	log.Printf("%s%s%s\n", green, message, reset)
}

func MyErrorLog(message string) {
	red := "\033[31m"
	reset := "\033[0m"
	log.Printf("%s%s%s\n", red, message, reset)
}

func MyInfoLog(message string) {
	blue := "\033[34m"
	reset := "\033[0m"
	log.Printf("%s%s%s\n", blue, message, reset)
}

func MyWarningLog(message string) {
	orange := "\033[33m"
	reset := "\033[0m"
	log.Printf("%s%s%s\n", orange, message, reset)
}

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
