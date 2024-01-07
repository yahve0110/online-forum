package routes

import (
	"fmt"
	"net/http"
	"strings"
	"forum/database" 
)

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered")
	if r.Method == http.MethodDelete {
		parts := strings.Split(r.URL.Path, "/")
		postID := parts[len(parts)-1]
		fmt.Println("At post")
		fmt.Println(postID)

	
		err := deletePostFromDB(postID)
		if err != nil {
			http.Error(w, "Failed to delete post", http.StatusInternalServerError)
			return
		}

	
		w.WriteHeader(http.StatusOK)
	} else {
		
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func deletePostFromDB(postID string) error {
	
	query := "DELETE FROM posts WHERE id = ?"


	_, err := database.Db.Exec(query, postID)
	if err != nil {
		return err
	}

	return nil
}
