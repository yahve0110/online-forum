package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/database"
	"forum/models"
	"log"

	"github.com/google/uuid"
)

// GenerateSessionToken создает уникальный сессионный токен
func GenerateSessionToken() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuidObj.String(), nil
}



// StoreSessionToken saves the session token with the user ID in the database

func StoreSessionToken(token, userID string) {
	_, err := database.Db.Exec("INSERT INTO sessions (token, user_id) VALUES (?, ?)", token, userID)
	if err != nil {
		log.Println("Error storing session token:", err)
	}
}
// GetUserByToken returns the user ID for a given session token from the database
func GetUserByToken(token string) (string, bool) {
	var userID string
	err := database.Db.QueryRow("SELECT user_id FROM sessions WHERE token = ?", token).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false
		}
		log.Println("Error getting user ID by token:", err)
		return "", false
	}
	return userID, true
}


// DeleteSessionToken deletes the session token from the database
func DeleteSessionsByUserID(userID string) {
	_, err := database.Db.Exec("DELETE FROM sessions WHERE user_id = ?", userID)
	if err != nil {
		log.Println("Error deleting sessions by user ID:", err)
	}
}




// ValidateSessionToken validates a session token and returns the associated user ID
 func ValidateSessionToken(token string) (string, bool) {
    userID, exists := GetUserByToken(token)
	fmt.Println(userID)
    return userID, exists
}


// GetUserByID retrieves user information by ID
func GetUserByID(userID string) (*models.User, error) {
	// Retrieve user data from the database (you should implement this function)
	var user models.User
	err := database.Db.QueryRow("SELECT id, nickname, age, gender, first_name, last_name, email, password FROM users WHERE id = ?", userID).
		Scan(&user.ID, &user.Nickname, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		log.Println("Error getting user by ID:", err)
		return nil, err
	}
	return &user, nil
}




// StoreUser saves user data in the database 
func StoreUser(user models.User) {
	_, err := database.Db.Exec("INSERT INTO users (id, username) VALUES (?, ?)", user.ID, user.Nickname)
	if err != nil {
		log.Println("Error storing user:", err)
	}
}