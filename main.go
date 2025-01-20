package main

import (
	"fmt"
	"net/http"
)
	var task1 = "Learn Go programming"
	var task2 = "Create a simple web application"
	var task3 = "Create a backend server"
	var tasks = []string{task1, task2, task3}
	
func main() {
    // fmt.Println("Hello, World!")
	http.HandleFunc("/",helloUser)
	http.HandleFunc("/showTasks", showTasks)

	http.ListenAndServe(":9090", nil)
	}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(writer, tasks)
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



// inorder create some datatype - make()
//FormatUint - converting uint to string from strconv package( it includes diff functions for convertings strings to and from other datatypes)
// time package - time.Sleep() - stops the execution of the thread(go routine) for given time.


//type Usersdata struct{    //creating struct using type(usertype data)
// 	firstName string
// 	lastName string
// 	email string
//  password string
// }



// GO concrrency  - 
//  with the use of 'go' keyword u can run required task in the background without disturbing the current execution flow.
// using let wg = sync.WaitGroup, "wg.Add(1)#1 - no. of tasks(go routines) before the go routine and wg.Wait()# at the end of main thread and wg.Done() at the end of the logic" we can wait for a go routinr to complete before the main routine is completed (as if the main routine is completed the backgrounf running task will be terminated so we need to wait until it gets completed)