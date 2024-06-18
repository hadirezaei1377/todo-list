package main

import (
	"fmt"
	"net/http"
	router "todo-list/routers"
)

func main() {
	r := router.NewRouter()

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
