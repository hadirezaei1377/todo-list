package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"todo-list/constant"
	"todo-list/contract"
	"todo-list/entity"
	"todo-list/repository/filestore"
	"todo-list/repository/memorystore"
	"todo-list/service/task"
)

var (
	userStorage     []entity.User
	categoryStorage []entity.Category

	authenticatedUser *entity.User
	serializationMode string
)

const (
	userStoragePath = "user.txt"
)

func main() {
	taskMemoryRepo := memorystore.NewTaskStore()

	taskService := task.NewService(taskMemoryRepo)

	serializeMode := flag.String("serialize-mode", constant.ManDarAvardiSerializationMode, "serialization mode to write data to file")
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	fmt.Println("Hello to TODO CLI app")

	switch *serializeMode {
	case constant.ManDarAvardiSerializationMode:
		serializationMode = constant.ManDarAvardiSerializationMode
	default:
		serializationMode = constant.JsonSerializationMode
	}

	var userFileStore = filestore.New(userStoragePath, serializationMode)

	// load user storage from file
	users := userFileStore.Load()
	userStorage = append(userStorage, users...)

	// if there is a user record with corresponding data allow the user to continue

	for {
		runCommand(userFileStore, *command, &taskService)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(store contract.UserWriteStore, command string, taskService *task.Service) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask(taskService)
	case "create-category":
		createCategory()
	case "register-user":
		registerUser(store)
	case "list-task":
		listTask(taskService)
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func createTask(taskService *task.Service) {

	scanner := bufio.NewScanner(os.Stdin)
	var title, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task category id")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category-id is not valid integer, %v\n", err)

		return
	}

	fmt.Println("please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	response, err := taskService.Create(task.CreateRequest{
		Title:               title,
		DueDate:             duedate,
		CategoryID:          categoryID,
		AuthenticatedUserID: authenticatedUser.ID,
	})

	if err != nil {
		fmt.Println("error", err)

		return
	}

	fmt.Println("create task:", response.Task)
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
	fmt.Println("category", title, color)

	c := entity.Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, c)
}

func registerUser(store contract.UserWriteStore) {
	scanner := bufio.NewScanner(os.Stdin)
	var id, name, email, password string

	fmt.Println("please enter the name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("user:", id, email, password)

	user := entity.User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: hashThePassword(password),
	}

	userStorage = append(userStorage, user)

	// writeUserToFile(user)
	store.Save(user)
}

func login() {
	fmt.Println("login process")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == hashThePassword(password) {
			authenticatedUser = &user
			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("the email or password is not correct")
	}
}

func listTask(taskService *task.Service) {
	userTasks, err := taskService.List(task.ListRequest{UserID: authenticatedUser.ID})
	if err != nil {
		fmt.Println("error", err)

		return
	}

	fmt.Println("user tasks", userTasks.Tasks)
}

func hashThePassword(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}
