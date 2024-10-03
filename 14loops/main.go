package main

import "fmt"

func main() {

	fmt.Println("Welcome to loops in golang")

	days := []string{"Sun", "Tue", "Wed"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and Day is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("Index is and Day is %v\n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		// if rougueValue == 5 {
		// 	break
		// }

		// if rougueValue == 5 {
		// 	rougueValue++
		// 	continue
		// }

		// if rougueValue == 5 {
		// 	rougueValue++
		// 	break
		// }

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("Jumping at LearnCodeOnline.co")
}
