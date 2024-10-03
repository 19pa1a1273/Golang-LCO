package main

import "fmt"

func main() {

	fmt.Println("Welcome to functions in golang")
	greeter()
	// greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is ", result)

	proRes, myMessage := proAdder(2, 5, 7, 8, 3)
	fmt.Println("Pro Result is ", proRes)
	fmt.Println("Pro Message is ", myMessage)
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"
}

// func greeterTwo() {
// 	fmt.Println("Another method")
// }
func greeter() {
	fmt.Println("Namastey in golnag")
}
