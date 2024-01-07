package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"forum/models"
	"forum/utils"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session token from the request cookies
	sessionCookie, err := r.Cookie("session_token")
	// Check if the session token is present
	if err == nil && sessionCookie != nil {
		// Validate the session token
		userID, valid := utils.ValidateSessionToken(sessionCookie.Value)

		// If the session token is valid, retrieve user ID from the database
		if valid {
			// Retrieve user information by user ID
			user, err := utils.GetUserByID(userID)
			if err == nil {
				// Retrieve all posts from the database
				posts, err := utils.GetAllPosts()
				if err == nil {
					// Respond with information indicating the user is logged in, including the User and Posts fields
					response := struct {
						IsLogged bool          `json:"isLogged"`
						User     models.User   `json:"user,omitempty"`
						Posts    []models.Post `json:"posts,omitempty"`
					}{
						IsLogged: true,
						User:     *user,
						Posts:    posts,
					}

					// Encode and send the response
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(response)
					return
				} else {
					log.Println("Error retrieving posts:", err)
				}
			} else {
				log.Println("Error retrieving user information:", err)
			}
		} else {
			log.Println("Invalid session token")
		}
	}

	// If no valid session token or an error occurred, respond with information indicating the user is not logged in
	response := struct {
		IsLogged bool `json:"isLogged"`
	}{
		IsLogged: false,
	}

	// Encode and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
