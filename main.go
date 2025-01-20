package main

import (
	"fmt"
)

//Global variables
var task1 = "Learn Go programming"
	var task2 = "Create a simple web application"
	var task3 = "Create a backend server"
	var tasks = []string{task1, task2, task3}
func main() {
    fmt.Println("Hello, World!")
	taskslist()
	}

func taskslist(){
	fmt.Println("List of Tasks:\n")
	for index, task := range tasks {
		fmt.Printf("Task %d: %s\n", index+1, task)
    
}
}