package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Web verb video - LCO")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {

	const myurl = "http://localhost:8000"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code is ", response.StatusCode)
	fmt.Println("Content Length is ", response.ContentLength)

	// In case we want to use the response of the body we will use ioutil.Readall()

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	// byteCount is just used to hold the count and we use responseString.Write(content) {Mainly Write() method}

	fmt.Println("Byte Count is ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform" : "learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Bhanu")
	data.Add("lastname", "Lingolu")
	data.Add("email", "bhanu@go.dev")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

// To call the localhost:8000 if lcowebserver is not there

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }

// func main() {
// 	http.HandleFunc("/", helloHandler)
// 	fmt.Println("Server is listening on port 8000...")
// 	err := http.ListenAndServe(":8000", nil)
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }
