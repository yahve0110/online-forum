package handlers

import (
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
   
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    switch r.URL.Path {
    case "/":
        MainPageHandler(w, r)
    case "/login":
        LoginHandler(w, r)
    case "/logout":
        logoutHandler(w, r)
    case "/register":
        RegisterHandler(w, r)
    default:
        // Handle unknown paths or return an error
        http.Error(w, "Not Found", http.StatusNotFound)
    }
}
