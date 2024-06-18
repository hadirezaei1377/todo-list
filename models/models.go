package models

type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
