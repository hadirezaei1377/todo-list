package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	Duedate    time.Time
	CategoryID int
	IsDone     bool
	UserID     uint
}

type Category struct {
	ID     int
	Title  string
	color  string
	UserID uint
}

var userStorage []User
var authenticatedUser *User
var taskStorage []Task
var categoryStorage []Category

func (u User) Print() {
	fmt.Println("user:", u.ID, u.Email, u.Name)
}

func main() {
	fmt.Println("welcome to your app!")

	command := flag.String("command", "no command", "command creates a new task from cli")
	flag.Parse()

	runCommand(*command)

}

func runCommand(command string) {

	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "list-task":
		ListTask()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid!", command)
	}
}

func createTask() {

	if authenticatedUser != nil {
		authenticatedUser.Print()
	}

	scanner := bufio.NewScanner(os.Stdin)
	var name, category, duedate string

	fmt.Println("please enter the task title")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the task category id")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println("category-id is not valid int, %v\n", err)

		return
	}

	isFound := false

	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == uint(authenticatedUser.ID) {

			break
		}
	}

	if !isFound {
		fmt.Println("category-id is not found, %v\n")

		return
	}

	fmt.Println("please enter the task duedate")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{

		ID:       len(taskStorage) + 1,
		Title:    title,
		Duedate:  duedate,
		Category: category,
		IsDone:   false,
		UserID:   uint(authenticatedUser.ID),
	}

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

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, category)

}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)

	var id, name, email, password string

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("user:", email, password)

	id = email

	fmt.Println("user:", id, email, password)

	user := User{
		ID:       rand.Int(),
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	// save user data in user.txt file, go create file
	path := "user.txt"
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("path doesnt exist", err)
		file, err := os.Create(path)

		return
		if err != nil {
			fmt.Println("cant create user.txt file", err)
		} else {

			file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY)
			if err != nil {
				fmt.Println("cant create or open file", err)

				return

			}

			data := fmt.Sprintf("id: %d, email: %s, name: %s, password: %s\n", user.ID, user.Email, user.Name,
				user.Password)

			file, err = os.Open(path)

			if err != nil {
				fmt.Println("path doesnt exist", err)

				return

			}

		}

		file.Write([]byte("new user record"))

		file.Close()

	}
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

	// get the email and password from the client
	// if there is a user record with corresponding data, allow user to continue
	fmt.Println("you must be logged in first")

	notFound := true

	if notFound {
		fmt.Println("the email or password is not correct")
		return
	}

	for _, user := range userStorage {
		if user.Email == email {
			if user.Password == password {
				notFound = false
				authenticatedUser = &user
				fmt.Println("you are logged in.")
			} else {
				fmt.Println("password is incorrect.")
			}
			break
		}
	}

	fmt.Println("user:", id, email, password)

}

func ListTask() {
	for _, task := range taskStorage {
		if task.UserID == uint(authenticatedUser.ID) {
			fmt.Println(task)
		}
	}
}
