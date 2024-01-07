package routes

import (
	"encoding/json"
	
	"forum/utils"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Credentials represents the JSON structure for login credentials
type Credentials struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

// LoginResponse represents the JSON structure for login response
type LoginResponse struct {
	IsLogged bool   `json:"isLogged"`
	Message  string `json:"message"`
	UserName string `json:"userName,omitempty"`
}

// LoginPageHandler handles requests to the "/login" route
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the method is correct
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON from the request
	var loginRequest Credentials
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate request fields
	if loginRequest.Identifier == "" || loginRequest.Password == "" {
		response := LoginResponse{
			IsLogged: false,
			Message:  "Missing required fields",
		}
		respondJSON(w, response, http.StatusBadRequest)
		return
	}

	// Get user by nickname or email
	user, err := utils.GetUserByIdentifier(loginRequest.Identifier)
	if err != nil {
		response := LoginResponse{
			IsLogged: false,
			Message:  "User not found",
		}
		respondJSON(w, response, http.StatusNotFound)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		response := LoginResponse{
			IsLogged: false,
			Message:  "Invalid credentials",
		}
		respondJSON(w, response, http.StatusUnauthorized)
		return
	}


	// If the password comparison is successful, generate a session token
	sessionToken, err := utils.GenerateSessionToken()
	if err != nil {
		response := LoginResponse{
			IsLogged: false,
			Message:  "Error generating session token",
		}
		respondJSON(w, response, http.StatusInternalServerError)
		return
	}
	



	// Set the session token as a cookie with HttpOnly and Secure attributes
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour), // Set the desired expiration time
		Path:     "/",    
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// Respond with a success message
	response := LoginResponse{
		IsLogged: true,
		Message:  "Login successful",
		UserName: user.Nickname, // Assuming UserName is the correct field
	}
	respondJSON(w, response, http.StatusOK)
		// Store the session token
		go utils.StoreSessionToken(sessionToken, user.ID)
}

func respondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}