package main

import "fmt"

func main(){
	var arr [5]int // Declare an array of 5 integers
	arr[0] = 10   // Assign value to the first element
	arr[1] = 20
	arr[2] = 30   // Assign value to the second element
	arr[3] = 40   // Assign value to the third element
	arr[4] = 50   // Assign value to the fourth element
	fmt.Println("Array elements:", arr) // Print the array elements
}
/*
array in golang is a collection of elements of the same type, stored in contiguous memory locations.
Arrays in Go have a fixed size, which is defined at the time of declaration.
Once an array is declared, its size cannot be changed.
*/