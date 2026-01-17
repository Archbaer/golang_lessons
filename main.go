package main

import (
	"fmt"
	"modulos/helper"
	"net/http"
)

var taskItems = []string{
	"Learn Go",
	"Build a web app",
	"Deploy the app",
}

func main() {
	fmt.Println("##### Welcome to my ToDo App")

	// http.HandleFunc("/", helloWorld)
	// http.HandleFunc("/show-tasks", showTasks)

	// http.ListenAndServe(":8080", nil)
	// var city string = "Istanbul"
	// fmt.Println("Enter city name:")
	// fmt.Scan(&city)

	// switch city {
	// case "London":
	// 	fmt.Println("The city is London")
	// case "New York":
	// 	fmt.Println("The city is New York")
	// case "Tokyo":
	// 	fmt.Println("The city is Tokyo")
	// }

	helper.PrintTasks(taskItems)

}

func showTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List of tasks:")
	for i, item := range taskItems {
		fmt.Fprintf(w, "%d. %s\n", i+1, item)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
