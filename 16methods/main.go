package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	bhanu := User{"bhanu", "bhanu@go.dev", true, 22}

	fmt.Println(bhanu)
	fmt.Printf("bhanu details are %+v\n", bhanu)
	fmt.Printf("Name is %v and Email is %v\n", bhanu.Name, bhanu.Email)
	bhanu.GetStatus()
	bhanu.NewEmail()
	fmt.Printf("Name is %v and Email is %v", bhanu.Name, bhanu.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
