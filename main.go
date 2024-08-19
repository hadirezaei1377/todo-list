package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	DueDate    string
	CategoryID int
	IsDone     bool
	UserID     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

var (
	userStorage     []User
	taskStorage     []Task
	categoryStorage []Category

	authenticatedUser *User
	serializationMode string
)

const (
	userStoragePath         = "user.txt"
	CustomSerializationMode = "custom"
	JsonSerializationMode   = "json"
)

var userFileStore = fileStore{
	filePath: userStoragePath,
}

func main() {

	serializeMode := flag.String("serialize-mode", CustomSerializationMode, "serialization mode to write data to file")
	command := flag.String("command", "no command", "command to run")
	flag.Parse()

	// load data from user storage file

	loadUserFromStorage(userFileStore, *serializeMode)

	fmt.Println("Hello to TODO app")

	switch *serializeMode {
	case CustomSerializationMode:
		serializationMode = CustomSerializationMode
	default:
		serializationMode = JsonSerializationMode
	}

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()
		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser(userFileStore)
	case "list-tasks":
		listTasks()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task category ID")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("Category ID is not a valid integer, %v\n", err)

		return
	}

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Printf("Category ID is not found.\n")

		return
	}

	fmt.Println("please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    duedate,
		CategoryID: categoryID,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)
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

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, category)
}

type userWriteStore interface {
	Save(u User)
}

type userReadStore interface {
	Load(serializationMode string) []User
}

func registerUser(store userWriteStore) {
	scanner := bufio.NewScanner(os.Stdin)

	var id, name, email, password string

	fmt.Println("please enter the user name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("user", id, name, email, password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: hashThePassword(password),
	}

	userStorage = append(userStorage, user)

	//writeUserToFile(user)
	store.Save(user)
}

func login() {
	fmt.Println("login process")

	scanner := bufio.NewScanner(os.Stdin)

	var email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	// get the email and password from the client

	for _, user := range userStorage {
		if user.Email == email && user.Password == hashThePassword(password) {
			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("The email or Password is not correct.")
	}

}

func listTasks() {
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}

func loadUserFromStorage(store userReadStore, serializationMode string) {
	users := store.Load(serializationMode)

	userStorage = append(userStorage, users...)
}

func (f fileStore) writeUserToFile(user User) {
	var file *os.File

	file, err := os.OpenFile(f.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("can't create or open file.", err)

		return
	}
	defer file.Close()

	var data []byte
	// serialize the user struct/object
	if serializationMode == CustomSerializationMode {
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", user.ID, user.Name, user.Email, user.Password))
	} else if serializationMode == JsonSerializationMode {
		// json
		var jErr error
		data, jErr = json.Marshal(user)
		if err != nil {
			fmt.Println("Can't marshal user struct to json.", jErr)

			return
		}

		data = append(data, []byte("\n")...)

	} else {
		fmt.Println("Invalid serialization mode.")

		return
	}

	numberOfWrittenBytes, wErr := file.Write(data)
	if wErr != nil {
		fmt.Println("can't write to the file %v\n", wErr)

		return
	}

	fmt.Println("numberOfWrittenBytes", numberOfWrittenBytes)
}

func deserializationFromCustom(userStr string) (User, error) {
	if userStr == "" {
		return User{}, errors.New("User string is empty.")
	}

	var user = User{}

	userFields := strings.Split(userStr, ",")
	for _, field := range userFields {
		values := strings.Split(field, ": ")
		if len(values) != 2 {
			fmt.Println("record is not valid", len(values))

			continue
		}
		fieldName := strings.ReplaceAll(values[0], " ", "")
		fieldValue := values[1]

		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				return User{}, errors.New("strconv error.")
			}
			user.ID = id
		case "name":
			user.Name = fieldValue
		case "email":
			user.Email = fieldValue
		case "password":
			user.Password = fieldValue
		}
	}

	return user, nil
}

func hashThePassword(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}

type fileStore struct {
	filePath string
}

func (f fileStore) Save(u User) {
	f.writeUserToFile(u)
}

func (f fileStore) Load(serializationMode string) []User {
	var uStore []User

	file, err := os.Open(f.filePath)
	if err != nil {
		fmt.Println("can't open the file", err)
	}

	var data = make([]byte, 1024)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("can't read from the file", oErr)

		return nil
	}

	var dataStr = string(data)

	userSlice := strings.Split(dataStr, "\n")

	for _, u := range userSlice {
		var userStruct = User{}

		switch serializationMode {
		case CustomSerializationMode:
			var dErr error
			userStruct, dErr = deserializationFromCustom(u)
			if dErr != nil {
				fmt.Println("Can't deserialize user record to user struct", dErr)

				return nil
			}
		case JsonSerializationMode:
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}
			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("Can't deserialize user record to user struct with json mode.", uErr)

				return nil
			}
		}

		userStorage = append(userStorage, userStruct)
	}

	return uStore
}
