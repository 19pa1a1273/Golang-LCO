package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to th time Study of golang")

	presentTime := time.Now() //time.Now().Nanosecond()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))           // we have to write as it is ("01-02-2006 15:04:05 Monday")
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Tuesday")) //It prints Tuesday only

	createdDate := time.Date(2022, time.March, 01, 06, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	// GOOS="linux" go build   -->executable file will be created
}
