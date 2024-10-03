package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	bhanu := User{"bhanu", "bhanu@go.dev", true, 22}

	fmt.Println(bhanu)
	fmt.Printf("bhanu details are %+v\n", bhanu)
	fmt.Printf("Name is %v and Email is %v", bhanu.Name, bhanu.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
