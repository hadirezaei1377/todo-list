package models

// User represents a user in the system
type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
}

// Todo represents a task to be done by a user
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
