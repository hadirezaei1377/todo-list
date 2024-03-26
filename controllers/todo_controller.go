package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/your-username/todo-list/db"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	var newTodo Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	// Save todo to Redis database
	db.SaveTodo(username, newTodo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	todos := db.GetTodos(username)
	json.NewEncoder(w).Encode(todos)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	todoID := params["id"]
	var updatedTodo Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)
	// Update todo in Redis database
	db.UpdateTodo(username, todoID, updatedTodo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	todoID := params["id"]
	// Delete todo from Redis database
	db.DeleteTodo(username, todoID)
	w.WriteHeader(http.StatusNoContent)
}
