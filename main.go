package main

import (
	"fmt"
	"modulos/greetings"
	"modulos/helper"
)

var taskItems = []string{
	"Learn Go",
	"Build a web app",
	"Deploy the app",
}

type UserData struct {
	Name     string
	Age      uint
	Email    string
	isMember bool
}

var user = UserData{
	Name:     "Alice",
	Age:      30,
	Email:    "alice.wonderland@outlook.com",
	isMember: true,
}

var userMap = map[string]UserData{
	"user1": {
		Name:     "Alice",
		Age:      30,
		Email:    "alice.wonderland@outlook.com",
		isMember: true,
	},
	"user2": {
		Name:     "Bob",
		Age:      25,
		Email:    "bob1.genko@gmail.com",
		isMember: false,
	},
}

func main() {
	fmt.Println("##### Welcome to my ToDo App")

	greetings.Greet("Alice")

	fmt.Printf("My favorite number is %d\n", greetings.MyNumber)

	helper.PrintTasks(taskItems)
	fmt.Printf("User Data: %+v\n", user)
	fmt.Printf("\n")
	fmt.Printf("User Map: %+v\n", userMap["user2"])

}
