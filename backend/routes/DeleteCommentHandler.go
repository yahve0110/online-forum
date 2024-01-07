package routes

import (
	
	"fmt"
	"net/http"
	"strings"
	"forum/database"
	_ "github.com/mattn/go-sqlite3"
)

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered DeleteCommentHandler")

	if r.Method == http.MethodDelete {
		// Extract the comment ID from the URL
		parts := strings.Split(r.URL.Path, "/")
		commentID := parts[len(parts)-1]
		fmt.Println("Comment ID:", commentID)

		// Your logic to delete the comment with commentID in the database
		err := deleteComment(commentID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Failed to delete comment")
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Comment deleted successfully")
	} else {
		// Method not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed")
	}
}

func deleteComment(commentID string) error {

	query := "DELETE FROM comments WHERE id = ?"
	_, err := database.Db.Exec(query, commentID)
	return err
}
