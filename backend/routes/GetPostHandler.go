package routes

import (
	"encoding/json"

	"forum/models"
	"forum/utils"
	"log"
	"net/http"
)

func GetPostDataHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the post ID from the URL
	postID := utils.ExtractPostID(r.URL.Path)

	// Get post information from the database
	post, err := utils.GetPostByID(postID)
	if err != nil {
		log.Println("Error retrieving post:", err)
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// Get comments for the post from the database
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




