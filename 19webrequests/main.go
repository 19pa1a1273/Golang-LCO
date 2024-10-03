package main

import (
	"fmt"
	"io"
	"net/http"
)

// const url = "https://lco.dev"
const url = "https://chaicode.com/"

func main() {

	fmt.Println("Lco web request")

	response, err := http.Get(url)

	checkError(err)
	fmt.Printf("Response is of type %T\n", response) //Response is of type *http.Response

	defer response.Body.Close() // Caller's responsibility to close the connection

	dataBytes, err := io.ReadAll(response.Body)

	checkError(err)

	content := string(dataBytes)
	fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
