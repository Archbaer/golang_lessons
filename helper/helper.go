package helper

import (
	"fmt"
)

func PrintTasks(arr []string) {
	for i, item := range arr {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}
