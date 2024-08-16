package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

const UserStoragepath = "user.txt"
const Defalse = "defalse"
const json = "json"

func (u User) Print() {
	fmt.Println("user:", u.ID, u.Email, u.Name)
}

func main() {

	// load user storage from file
	loadUserStorageFromFile(*serializationMode)

	fmt.Println("welcome to your app!")

	serilizeMode := flag.String("serialize mode", "defalse", "serialization mode to write data to file")

	command := flag.String("command", "no command", "command creates a new task from cli")
	flag.Parse()

	switch *serilizeMode {
	case "defalse":
		serializationMode = Defalse
	default:

		serializationMode = json

	}

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

	WriteUserToFile(user)

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

func loadUserStorageFromFile(serializationMode string) {

	file, err := os.Open(UserStoragepath)

	if err != nil {
		fmt.Println("cant open the file", err)

		return

	}

	var data = make([]byte, 10240)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("cant read from the file", err)

		return

	}

	var dataStr = string(data)

	userSlice := strings.Split(dataStr, "\n")
	fmt.Println("userSlice:", len(userSlice))
	for _, user := range userSlice {
		var userStruct = User{}
		switch serializationMode {
		case Defalse:
			var dErr error
			userStruct, dErr = deseializeFromDefalse(user)

			if dErr != nil {
				fmt.Println("cant deserialize user record to user struct", err)

				return
			}
		case json:

			uErr := json.Unmarshal([]byte(user), &userStruct)

			if uErr != nil {
				fmt.Println("cant deserialize user record to user struct with json mode", uErr)

				return

			}
		}
		fmt.Println("unmarshalled user", userStruct)
		userStorage := append(userStorage, userStruct)

	}

}

func WriteUserToFile(user User) {

	var file *os.File

	file, err := os.OpenFile(UserStoragepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cant create or open file", err)

		return

	}
	defer file.Close()

	var data []byte

	if serializationMode == Defalse {
		data = []byte(fmt.Sprintf("id: %d, email: %s, name: %s, password: %s\n", user.ID, user.Email, user.Name,
			user.Password))

	} else if serializationMode == json {

		data, err = json.Marshal(user)

		if err != nil {
			fmt.Println("cant marshal user struct to json", err)

			return
		}

	} else {
		fmt.Println("invalid serializatiion mode")

		return
	}

	var b = data

	numberOfWrittenBytes, wErr := file.Write(b)
	if wErr != nil {
		fmt.Println("cant write to the file %v\n ", wErr)

		return

	}

	fmt.Println("numberOfWrittenBytes: ", numberOfWrittenBytes)

}

func deseializeFromDefalse(userStr string) (User, error) {
	if userStr == "" {

		return User{}, errors.New("user string is empty")
	}

	fmt.Println("line of file:", index, "data:", data)

	var user = User{}

	userFields := strings.Split(userStr, ",")

	for _, field := range userFields {
		fmt.Println(field)

		values := strings.Split(field, ": ")
		fieldName := strings.ReplaceAll(values[0], " ", "")
		fieldValue := values[1]

		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)

			if err != nil {
				fmt.Println(err)

				return User{}, errors.New("str conv err")

			}

			user.ID = fieldValue
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
