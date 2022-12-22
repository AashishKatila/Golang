package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Package")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //To give present date time and day

	createdDate := time.Date(2025, time.July, 6, 23, 33, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
