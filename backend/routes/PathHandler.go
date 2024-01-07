package routes

import (
	"fmt"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	
	switch {
	case r.URL.Path == "/":
		MainPageHandler(w, r)
	case r.URL.Path == "/login":
		LoginPageHandler(w, r)
	case r.URL.Path == "/register":
		RegisterPageHandler(w, r)
	case r.URL.Path == "/logout":
		LogoutHandler(w, r)
	case r.URL.Path == "/create-post":
		CreatePostHandler(w, r)
	case strings.HasPrefix(r.URL.Path, "/delete-post/"):
		DeletePostHandler(w, r)
	case strings.HasPrefix(r.URL.Path, "/delete-comment/"):
		DeleteCommentHandler(w, r)
	default:
		// Check if the path starts with "/add-comment"
		if strings.HasPrefix(r.URL.Path, "/add-comment") {
			AddCommentHandler(w, r)
			return
		}
		// Check if the path starts with "/post/"
		if strings.HasPrefix(r.URL.Path, "/post/") {
			GetPostDataHandler(w, r)
		} else {
			RedirectHome(w, r)
		}
	}
}


func RedirectHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirecting to home page")
	http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
}
