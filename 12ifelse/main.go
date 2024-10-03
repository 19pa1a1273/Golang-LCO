package main

import "fmt"

func main() {

	fmt.Println("If else in golang")

	loginCount := 20
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login Count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if n := 3; n < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("NUm is NOT less than 10")
	}

	// if err != nill{

	// }
}
