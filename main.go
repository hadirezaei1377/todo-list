package main

import (
	"fmt"
	"net/http"
	"todo-list/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{username}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{username}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{username}", controllers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/todos/{username}", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{username}", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{username}/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{username}/{id}", controllers.DeleteTodo).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
