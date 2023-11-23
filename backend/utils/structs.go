package utils

type User struct{
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Age       string    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Session   string `json:"session"`
}