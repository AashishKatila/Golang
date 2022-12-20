package main

import "fmt"

func main() {

	//Declare maps and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "Sharon@gmail.com"}
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

}
