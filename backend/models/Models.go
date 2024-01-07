package models

import "time"


type User struct{
	ID 	string `json:"id"`
	Nickname string `json:"nickname"`
	Age string `json:"age"`
	Gender string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var LoginData struct{
	Identifier    string `json:"identifier"`
	Password string `json:"password"`
}

type Post struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID           string    `json:"id"`
	PostID       string    `json:"post_id"`
	Author       string    `json:"author"`
	Content      string    `json:"content"`
	CreationTime time.Time `json:"creation_time"`
}

type PostWithComments struct {
	Post     *Post  `json:"post"`
	Comments []Comment     `json:"comments"`
}
