package routes

import (
	"forum/utils" // Assuming this is your utility package
	"net/http"
	"time"
)

// LogoutPageHandler handles requests to log out
// LogoutPageHandler handles requests to log out
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session token cookie
	sessionCookie, err := r.Cookie("session_token")
	if err != nil {
		// If the cookie is not found, the user is already logged out
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Already logged out"))
		return
	}

	// Delete the session token from the server-side storage
	utils.DeleteSessionsByUserID(sessionCookie.Value)

	// Set a new session token cookie with an expired expiration time to invalidate the existing session on the client side
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Set expiration in the past
		Path:     "/",                        // Set the same path as the original cookie
		HttpOnly: true,
		Secure:   true, // Set to true if using HTTPS
        SameSite: http.SameSiteNoneMode,
	})

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
