package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/database"
	"forum/utils"
	"log"
	"net/http"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit register")

	//only post methods allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming JSON payload

    var newUser utils.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
        return
    }

    fmt.Println("Decoded user data:", newUser)
	
	if newUser.Nickname == "" || newUser.Age == "" || newUser.Email == "" || newUser.Password == "" || newUser.Gender == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if the user already exists by nickname or email
	existingUser, err := GetUserByNicknameOrEmail(newUser.Nickname, newUser.Email)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Error checking if user exists", http.StatusInternalServerError)
		return
	}
	if existingUser.ID != "" {
		http.Error(w, "User with this nickname or email already exists", http.StatusConflict)
		return
	}

	// Generate a unique ID for the new user
	newUser.ID = uuid.New().String()
	fmt.Println("Decoded user data:", newUser)

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	newUser.Password = string(hashedPassword)

	// Insert the user into the database
	
	err = InsertUser(newUser)
	if err != nil {
		http.Error(w, "Error inserting user into the database", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "User registered successfully", "id": newUser.ID}
	json.NewEncoder(w).Encode(response)
}

// GetUserByNicknameOrEmail retrieves a user by nickname or email
func GetUserByNicknameOrEmail(nickname, email string) (utils.User, error) {
	var user utils.User
	err := database.Db.QueryRow(`
		SELECT id, nickname, age, first_name, last_name, email, password
		FROM users
		WHERE nickname = ? OR email = ?
	`, nickname, email).Scan(&user.ID, &user.Nickname, &user.Age, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	return user, err
}

// InsertUser inserts a new user into the database
func InsertUser(user utils.User) error {

   _, err := database.Db.Exec(`
   INSERT INTO users (id, nickname, age, gender, first_name, last_name, email, password)
   VALUES (?, ?, ?, ?, ?, ?, ?, ?)
  `, user.ID, user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
  
  if err != nil {
	  log.Println("Error inserting user into the database:", err)
	  return fmt.Errorf("error inserting user into the database: %v", err)
  }
  return err
}
