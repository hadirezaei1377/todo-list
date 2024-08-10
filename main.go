package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

type User struct {
	ID       int
	email    string
	password string
}

var userStorage []User

func main() {
	fmt.Println("welcome to your app!")

	command := flag.String("command", "no command", "command creates a new task from cli")
	flag.Parse()

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

	fmt.Printf("userStorage: %+v\n", userStorage)
}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid!", command)
	}
}

func createTask() {

	scanner := bufio.NewScanner(os.Stdin)
	var name, category, duedate string

	fmt.Println("please enter the task title")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("please enter the task duedate")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("task:", name, category, duedate)

}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)

	var title, color string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("category:", title, color)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)

	var id, email, password string

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("user:", email, password)

	id = email

	fmt.Println("user:", id, email, password)

	user := User{
		ID:       rand.Int(),
		email:    email,
		password: password,
	}

	userStorage = append(userStorage, user)

}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("user:", id, email, password)
}
