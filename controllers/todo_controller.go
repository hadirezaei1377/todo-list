package controllers

import (
	"encoding/json"
	"net/http"
	"todo-list/db"

	"github.com/gorilla/mux"
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
	db.UpdateTodo(username, todoID, updatedTodo)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	todoID := params["id"]
	db.DeleteTodo(username, todoID)
	w.WriteHeader(http.StatusNoContent)
}
