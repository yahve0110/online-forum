package routes

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"forum/utils"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type PostData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: Create post")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the session token from the request cookies
	sessionCookie, err := r.Cookie("session_token")

	
	// Check if the session token is present
	if err != nil || sessionCookie == nil {
		http.Error(w, "Invalid or missing session token", http.StatusUnauthorized)
		return
	}

	// Validate the session token
	userID, valid := utils.ValidateSessionToken(sessionCookie.Value)
	fmt.Println(userID)
	// If the session token is valid, retrieve user ID from the database
	if !valid {
		log.Println("Invalid session token")
		http.Error(w, "Invalid session token", http.StatusUnauthorized)
		return
	}

	// Retrieve user information by user ID
	user, err := utils.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting user by ID", http.StatusInternalServerError)
		return
	}

	// Generate post ID
	postID := uuid.New().String()
	// Set the current time as the post creation time
	creationTime := time.Now()

	var postData PostData
	if err := json.NewDecoder(r.Body).Decode(&postData); err != nil {
		http.Error(w, "Error while parsing JSON", http.StatusBadRequest)
		return
	}

	// Insert the new post into the database
	_, err = database.Db.Exec(`
		INSERT INTO posts (id, user_nickname, title, content, creation_time)
		VALUES (?, ?, ?, ?, ?);
	`, postID, user.Nickname, postData.Title, postData.Content, creationTime)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error inserting post into the database", http.StatusInternalServerError)
		return
	}

	// Return a successful response
	response := map[string]interface{}{
		"status":   "success",
		"message":  "Post created successfully",
		"post_id":  postID,
		"username": user.Nickname,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
