package db

import (
	"encoding/json"
	"fmt"
	"todo-list/models"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis!")
}

func SaveUser(user models.User) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	err = client.Set(user.Username, userJSON, 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetUser(username string) models.User {
	val, err := client.Get(username).Result()
	if err != nil {
		panic(err)
	}
	var user models.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		panic(err)
	}
	return user
}

func UpdateUser(username string, updatedUser models.User) {
	DeleteUser(username)
	SaveUser(updatedUser)
}

func DeleteUser(username string) {
	err := client.Del(username).Err()
	if err != nil {
		panic(err)
	}
}

func SaveTodo(username string, todo models.Todo) {
	todoJSON, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	err = client.HSet(username, todo.ID, todoJSON).Err()

	if err != nil {
		panic(err)
	}
}

func GetTodos(username string) []models.Todo {
	val, err := client.HGetAll(username).Result()
	if err != nil {
		panic(err)
	}
	todos := make([]models.Todo, len(val))
	i := 0
	for _, v := range val {
		var todo models.Todo
		err := json.Unmarshal([]byte(v), &todo)
		if err != nil {
			panic(err)
		}
		todos[i] = todo
		i++
	}
	return todos
}

func UpdateTodo(username string, todoID string, updatedTodo models.Todo) {
	DeleteTodo(username, todoID)
	SaveTodo(username, updatedTodo)
}

func DeleteTodo(username string, todoID string) {
	err := client.HDel(username, todoID).Err()
	if err != nil {
		panic(err)
	}
}
