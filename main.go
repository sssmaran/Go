package main

import (
	"fmt"
	"net/http"
)

func main() {
    // fmt.Println("Hello, World!")
	http.HandleFunc("/",helloUser)

	http.ListenAndServe(":9090", nil)
	// //local variables
	// var task1 = "Learn Go programming"
	// var task2 = "Create a simple web application"
	// var task3 = "Create a backend server"
	// var tasks = []string{task1, task2, task3}
	// tasksList(tasks)
	// fmt.Println()
	// addTask(tasks, "Use HTMX for frontend design")
	}

func helloUser(writer http.ResponseWriter, request *http.Request){
	var greeting = "Hello, Welcome to Golang"
	fmt.Fprintln(writer, greeting)
}

func tasksList(tasks []string) {       //parameter passinng tasks list
	fmt.Println("List of Tasks:")
	for index, task := range tasks {
		fmt.Printf("Task %d: %s\n", index+1, task)
    
}
}

func addTask(tasks []string, newTask string) {
	var updatedTasks = append(tasks, newTask)
    tasksList(updatedTasks)
}