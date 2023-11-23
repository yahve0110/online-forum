package main

import (
	"fmt"
	"forum/database"
	"forum/handlers"
	"net/http"
    "forum/utils"
)

func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
        w.Header().Set("Access-Control-Allow-Credentials", "true")


        // Allow preflight requests to proceed without further processing
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Call the next handler in the chain
        next.ServeHTTP(w, r)
    })
}



func main() {
	database.InitDB()
	defer database.CloseDB()

	go utils.CleanUpSessionsScheduler()

	// Use the CORS middleware for all routes
	http.Handle("/", CORSMiddleware(http.HandlerFunc(handlers.PathHandler)))

	fmt.Println("Starting server at port 8080")
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
