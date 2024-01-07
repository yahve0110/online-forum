package utils

import (
	"database/sql"
	"fmt"
	"forum/database"
	"forum/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//validate data
func ValidateUserData(w http.ResponseWriter, newUser models.User) bool {
	if newUser.Nickname == "" || newUser.Age == "" || newUser.Gender == "" || newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" || newUser.Password == "" {
		errorMessage := "Required fields are missing: "
		if newUser.Nickname == "" {
			errorMessage += "Nickname "
		}
		if newUser.Age == "" {
			errorMessage += "Age "
		} else {
			_, err := strconv.Atoi(newUser.Age)
			if err != nil {
				errorMessage += "Invalid Age format "
			}
		}
		if newUser.Gender == "" {
			errorMessage += "Gender "
		}
		if newUser.FirstName == "" {
			errorMessage += "First Name "
		}
		if newUser.LastName == "" {
			errorMessage += "Last Name "
		}
		if newUser.Email == "" {
			errorMessage += "Email "
		}
		if newUser.Password == "" {
			errorMessage += "Password "
		}
		http.Error(w, `{"error": "`+errorMessage+`"}`, http.StatusBadRequest)
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// GetUserByIdentifier retrieves a user by either nickname or email
func GetUserByIdentifier(identifier string) (*models.User, error) {
	var user models.User

	// Query the database to retrieve the user by nickname or email
	err := database.Db.QueryRow(`
		SELECT id, nickname, age, gender, first_name, last_name, email, password
		FROM users
		WHERE nickname = ? OR email = ?
	`, identifier, identifier).Scan(
		&user.ID,
		&user.Nickname,
		&user.Age,
		&user.Gender,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// No user found with the given identifier
			return nil, nil
		}
		log.Println("Error retrieving user:", err)
		return nil, err
	}

	return &user, nil
}

// GetAllPosts retrieves all posts from the database
func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post

	rows, err := database.Db.Query("SELECT * FROM posts")
	if err != nil {
		log.Println("Error retrieving posts:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning post:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}

	return posts, nil
}



func ExtractPostID(path string) string {
	// Extract the post ID from the path 

	segments := strings.Split(path, "/")
	return segments[len(segments)-1]
}

// GetPostByID retrieves information about a post by its ID
func GetPostByID(postID string) (*models.Post, error) {
	var post models.Post

	// Query the database to get post information
	err := database.Db.QueryRow("SELECT * FROM posts WHERE id = ?", postID).Scan(
		&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post not found")
		}
		return nil, err
	}

	return &post, nil
}

func GetCommentsForPost(postID string) ([]models.Comment, error) {
	rows, err := database.Db.Query("SELECT * FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return []models.Comment{}, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.Author, &comment.Content, &comment.CreationTime)
		if err != nil {
			return []models.Comment{}, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}