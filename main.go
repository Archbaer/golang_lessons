package main

import (
	"fmt"
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
	intChannel := make(chan int)

	go func() {
		intChannel <- 2
	}()

	fmt.Println(<-intChannel)
}
