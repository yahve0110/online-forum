package routes

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"forum/models"
	"forum/utils"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)



func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	
	// Extract the post ID from the URL
	postID := utils.ExtractPostID(r.URL.Path)

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
	// commentID := uuid.New().String()
	// Set the current time as the post creation time
	creationTime := time.Now()

// Extract comment data from the form
author := user.Nickname // Assuming you want to use the user's nickname as the comment author
   // Read the request body
   body, err := ioutil.ReadAll(r.Body)
   if err != nil {
	   log.Println("Error reading request body:", err)
	   http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	   return
   }

   // Unmarshal JSON data
   var requestData map[string]string
   if err := json.Unmarshal(body, &requestData); err != nil {
	   log.Println("Error unmarshalling JSON:", err)
	   http.Error(w, "Bad Request", http.StatusBadRequest)
	   return
   }

   // Get the content from the JSON data
   content := requestData["content"]
   fmt.Printf("Received comment content: %s\n", content)

// Perform validation on comment content
if content == "" {
    http.Error(w, "Comment content is required", http.StatusBadRequest)
    return
}

err = AddCommentToPost(postID, author, content, creationTime)
if err != nil {
    log.Println("Error adding comment:", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
}


	// Get updated post and comments information
	post, err := utils.GetPostByID(postID)
	if err != nil {
		log.Println("Error retrieving post:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	comments, err := utils.GetCommentsForPost(postID)
	if err != nil {
		log.Println("Error retrieving comments:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Combine post and comments information
	postWithComments := models.PostWithComments{
		Post:     post,
		Comments: comments,
	}

	// Convert the combined information to JSON
	postWithCommentsJSON, err := json.Marshal(postWithComments)
	if err != nil {
		log.Println("Error marshalling post with comments to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the content type and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(postWithCommentsJSON)
}

// AddCommentToPost добавляет комментарий к посту в базе данных
func AddCommentToPost(postID, author, content string, creationTime time.Time) error {
	_, err := database.Db.Exec(`
		INSERT INTO comments (id, post_id, author, content, creation_time)
		VALUES (?, ?, ?, ?, ?)
	`, uuid.New().String(), postID, author, content, creationTime)

	if err != nil {
		log.Println("Error adding comment to post:", err)
		return err
	}

	log.Println("Comment added to post successfully")
	return nil
}
