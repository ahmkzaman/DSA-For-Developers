package main

import (
	"fmt"
)

/*
func main(){
	var arr [3]int // Declare an array of 5 integers
	arr[0] = 10   // Assign value to the first element
	arr[1] = 20
	arr[2] = 30   // Assign value to the second element
	fmt.Println("Array elements:", arr) // Print the array elements
}
*/

/*
array in golang is a collection of elements of the same type, stored in contiguous memory locations.
Arrays in Go have a fixed size, which is defined at the time of declaration.
Once an array is declared, its size cannot be changed.
*/

//Real world example
type Task struct {
	ID    int
	Title string
	Complete bool
}

func main(){
	tasks := [3]Task{
		{ID: 1, Title: "Task 1", Complete: false},
		{ID: 2, Title: "Task 2", Complete: true},
		{ID: 3, Title: "Task 3", Complete: false},
			
	}
	//Print the tasks

	fmt.Println("Tasks:")
	for _, task := range tasks {
	fmt.Printf("ID: %d, Title: %s, Complete: %t\n", task.ID, task.Title, task.Complete)
}

}
