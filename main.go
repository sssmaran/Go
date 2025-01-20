package main

import (
	"fmt"
)

func main() {
    fmt.Println("Hello, World!")
	//local variables
	var task1 = "Learn Go programming"
	var task2 = "Create a simple web application"
	var task3 = "Create a backend server"
	var tasks = []string{task1, task2, task3}
	taskslist(tasks)
	}

func taskslist(tasks []string) {       //parameter passinng tasks list
	fmt.Println("List of Tasks:")
	for index, task := range tasks {
		fmt.Printf("Task %d: %s\n", index+1, task)
    
}
}