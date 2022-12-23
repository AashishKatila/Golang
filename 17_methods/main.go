package main

import "fmt"

func main() {
	john := User{"John", "john@abcd.com", "Active", 19}
	fmt.Println("Name is", john.Name)
	john.getStatus()
	john.getEmail()
	fmt.Println(john)
}

type User struct {
	Name   string
	Email  string
	Status string
	Age    int
}

func (u User) getStatus() {
	fmt.Println("User is", u.Status)
}

func (u User) getEmail() {
	u.Email = "email@xyz.com"
	fmt.Println("User is", u.Email)
}
