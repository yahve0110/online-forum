package handlers

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"forum/utils"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct{
	Identifier string `json:"identifier`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: login")
	
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//decoding JSON from request
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil{
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
        return
	}

	//validation of request fields
	if loginRequest.Identifier == "" || loginRequest.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	//get user by nickname or email
	user,err := GetUserByIdentifier(loginRequest.Identifier)
	if err != nil{
		http.Error(w, "Error checking if user exists", http.StatusInternalServerError)
		return
	}

	//compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//create new session
	session := utils.NewSession()
	session.Values["user"] = user

	//save session
	utils.SetSession(session)

	//set session id in cookie

	cookie := http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Expires:  time.Now().Add(24 * time.Hour), // 24 hours 
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
http.SetCookie(w, &cookie)

	//send response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Login successful"}
	json.NewEncoder(w).Encode(response)
}


func GetUserByIdentifier(identifier string) (utils.User, error) {
	var user utils.User
	err := database.Db.QueryRow(`
		SELECT id, nickname, age, gender, first_name, last_name, email, password
		FROM users
		WHERE nickname = ? OR email = ?
	`, identifier, identifier).Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	return user, err
}

