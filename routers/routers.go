package router

import (
	"net/http"
	"todo-list/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	// Users endpoints
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{username}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{username}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{username}", controllers.DeleteUser).Methods("DELETE")

	// Todos endpoints
	r.HandleFunc("/todos/{username}", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{username}", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{username}/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{username}/{id}", controllers.DeleteTodo).Methods("DELETE")

	return r
}
