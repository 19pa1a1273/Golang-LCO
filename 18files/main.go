package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - LCO.in"
	file, err := os.Create("./mylcogofile")

	// if err != nil {
	// 	panic(err)
	// }

	chechNilErr(err)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	chechNilErr(err)

	fmt.Println("Length is ", length)
	defer file.Close()
	readFile("./mylcogofile")

}

func readFile(filename string) {

	// dataByte, err := ioutil.ReadFile(filename)
	dataByte, err := os.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }
	chechNilErr(err)

	// fmt.Println("Text inside the text file is \n", dataByte) //output is in values like 84 104 105 ....
	fmt.Println("Text inside the text file is \n", string(dataByte))

}

func chechNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
