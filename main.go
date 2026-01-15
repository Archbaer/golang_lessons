package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{
	"Learn Go",
	"Build a web app",
	"Deploy the app",
}

func main() {
	fmt.Println("##### Welcome to my ToDo App")

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List of tasks:")
	for i, item := range taskItems {
		fmt.Fprintf(w, "%d. %s\n", i+1, item)
	}
}

func printTasks(arr []string) {
	for i, item := range arr {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// func addTasks()
