package main

import "fmt"

const LoginToken string = "abcdef" //public

//jwtToken := 300000 //This is not allowed outside of Method
// var jwtToken int = 300000

func main() {
	var username string = "Bhanu"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.35356487867345635
	fmt.Println(smallFloat) //255.35356
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var smallFloat64 float64 = 255.35356487867345635
	fmt.Println(smallFloat64) //255.35356487867347
	fmt.Printf("Variable is of type: %T \n", smallFloat64)

	//default value and some aliases

	var anotherInteger int
	fmt.Println(anotherInteger)
	fmt.Printf("Variable is of type: %T \n", anotherInteger)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	//implicit type

	var website = "learning.com"
	fmt.Println(website)

	//no var type

	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	//using const variable

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
