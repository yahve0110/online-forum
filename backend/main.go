package main

import (
	"fmt"
	"forum/database"
	"forum/routes"
	"net/http"
)


func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers dynamically based on the request's Origin header
        origin := r.Header.Get("Origin")
        if origin != "" {
            w.Header().Set("Access-Control-Allow-Origin", origin)
            w.Header().Set("Access-Control-Allow-Credentials", "true")
        }

        // Allow only specific methods for actual requests
        if r.Method == http.MethodOptions {
            w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
            w.WriteHeader(http.StatusOK)
            return
        }

        // Call the next handler in the chain for actual requests
        next.ServeHTTP(w, r)
    })
}

func main() {
	database.InitDB()
	defer database.Db.Close()

	// Create a new ServeMux for routing
	router := http.NewServeMux()

	// Add your routes here
	router.HandleFunc("/", routes.Router)

	// Wrap the router with CORS middleware
	handler := CORSMiddleware(router)

	fmt.Println("Starting server at port 8080")
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", handler)
}
