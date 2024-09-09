package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/yaviral17/hw-go/db"
	"github.com/yaviral17/hw-go/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// Parse the incoming JSON request
	var userLogin models.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	// Validate the credentials against the database
	ctx := context.Background()
	user, err := db.GetUserByLoginCredentials(ctx, userLogin)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Invalid username or password",
			"error":   err.Error(),
		})
		return
	}

	// Return the user details as a JSON response also add status and message
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "User logged in successfully",
		"user":    user,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding request body: ", err)
		// send a json response with status and message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx := context.Background()
	uid, err := db.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error creating user: ", err)
		// send a json response with status and message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// Return the user details as a JSON response also add status and message
	(w).Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "User registered successfully 🎉",
		"user_id": uid,
	})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
