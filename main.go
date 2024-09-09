package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yaviral17/hw-go/db"
	"github.com/yaviral17/hw-go/myLogs"
)

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
	loggedRouter := myLogs.LoggingMiddleware(router)

	http.ListenAndServe("0.0.0.0:"+port, loggedRouter) // Start the server
}
