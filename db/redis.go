package db

import (
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "your_redis_password", // set your Redis password here
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis!")
}

func SaveUser(user User) {
	// Save user to Redis
}

func GetUser(username string) User {
	// Get user from Redis
}

func UpdateUser(username string, updatedUser User) {
	// Update user in Redis
}

func DeleteUser(username string) {
	// Delete user from Redis
}

func SaveTodo(username string, todo Todo) {
	// Save todo to Redis
}

func GetTodos(username string) []Todo {
	// Get todos from Redis
}

func UpdateTodo(username string, todoID string, updatedTodo Todo) {
	// Update todo in Redis
}

func DeleteTodo(username string, todoID string) {
	// Delete todo from Redis
}