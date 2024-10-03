package main

import "fmt"

func main() {

	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all  languages: ", languages)

	delete(languages, "RB")

	fmt.Println("List of all  languages: ", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For Key , Value is %v\n", value)
	}

}
