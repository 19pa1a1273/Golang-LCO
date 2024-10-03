package main

import "fmt"

func main() {

	fmt.Println("Welcome to a class on Pointers")

	// var ptr *int
	// fmt.Println("Value of a pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber // assigning and referencing the pointer
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is: ", myNumber)
}
