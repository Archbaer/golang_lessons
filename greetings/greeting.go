package greetings

import (
	"fmt"
)

var MyNumber = 67

func Greet(name string) {
	if name == "" {
		fmt.Println("Hello, Guest!")
	} else {
		fmt.Printf("Hello, %v! \n", name)
	}
}
