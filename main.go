package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to your app!")

	command := flag.String("command", "no command", "command creates a new task from cli")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	if *command == "create-task" {
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

	} else if *command == "create category" {
		var title, color string

		fmt.Println("please enter the category title")
		scanner.Scan()
		title = scanner.Text()

		fmt.Println("please enter the category color")
		scanner.Scan()
		color = scanner.Text()

		fmt.Println("category:", title, color)

	} else if *command == "register-user" {

		var id, email, password string

		fmt.Println("please enter the category email")
		scanner.Scan()
		email = scanner.Text()

		fmt.Println("please enter the category password")
		scanner.Scan()
		password = scanner.Text()

		fmt.Println("category:", email, password)

		id = email

		fmt.Println("user:", id, email, password)
	}
}
