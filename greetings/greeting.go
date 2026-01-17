package greetings

import (
	"fmt"
)

func Greet(name string) {
	if name == "" {
		fmt.Println("Hello, Guest!")
	} else {
		fmt.Printf("Hello, %v! \n", name)
	}
}
