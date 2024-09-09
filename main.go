package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yaviral17/hw-go/db"
	"github.com/yaviral17/hw-go/myLogs"
	"github.com/yaviral17/hw-go/routes"
)

func main() {

	// if godotenv.Load() != nil {
	// 	myLogs.MyErrorLog("Error loading .env file")
	// 	// return
	// }// Conditionally load .env file only in development
	if os.Getenv("ENV") == "development" {
		if err := godotenv.Load(); err != nil {
			myLogs.MyErrorLog("Error loading .env file")
			return
		}
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println("DATABASE_URL: ", dbUrl)
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

	// auth routes
	apiRouter.HandleFunc("/auth/login", routes.Login).Methods("POST")
	apiRouter.HandleFunc("/auth/register", routes.Register).Methods("POST")

	// Apply the logging middleware to the router
	loggedRouter := myLogs.LoggingMiddleware(router)
	log.Println("Server running on port: ", port)

	http.ListenAndServe("0.0.0.0:"+port, loggedRouter) // Start the server
}
