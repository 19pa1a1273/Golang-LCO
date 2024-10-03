package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two") // LIFO
	fmt.Println("Hello")
	myDefer()
}

// World, One, Two // LIFO
// 0, 1, 2, 3, 4
// hello, 43210, Two, One, World

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
