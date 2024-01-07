package routes

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"forum/models"
	"forum/utils"
	"net/http"
	"github.com/google/uuid"
)


func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: Registration page")

	//check if method is correct
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//variable to store data from request
	var newUser models.User

	userID := uuid.New().String()

	newUser.ID = userID
	//parse data 
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		http.Error(w, `{"error": "Error while parsing JSON"}`, http.StatusBadRequest)
		return
	}
	
	fmt.Printf("Received JSON data: %+v\n", newUser)


	// Validate data
	if !utils.ValidateUserData(w, newUser) {
		fmt.Println("Error validating user data")
		return
	}
	

	//hash password 
	hashedPassword,err := utils.HashPassword(newUser.Password)
	if err != nil{
		http.Error(w, `{"error": "Error while hashing Password"}`, http.StatusInternalServerError)
		return
	}

	//change password to hashed
	newUser.Password = hashedPassword


	// Check if user with the same nickname already exists
	var count int
	err = database.Db.QueryRow("SELECT COUNT(*) FROM users WHERE nickname = ?", newUser.Nickname).Scan(&count)
	if err != nil {
		http.Error(w, `{"error": "Error checking user existence"}`, http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, `{"error": "User with this nickname already exists"}`, http.StatusBadRequest)
		return
	}

	// Insert the new user into the database
	_, err = database.Db.Exec(`
    INSERT INTO users (id, nickname, age, gender, first_name, last_name, email, password)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?);
`, newUser.ID, newUser.Nickname, newUser.Age, newUser.Gender, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password)

	if err != nil {
		http.Error(w, `{"error": "Error inserting user into database"}`, http.StatusInternalServerError)
		return
	}

    fmt.Printf("New User: %s\n", newUser)

	response := map[string]interface{}{
		"status": "success",
		"message": "User registered successfully",
		"user": newUser,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(response)

}
