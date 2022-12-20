package main

import (
	"fmt"
	"strconv"
)

// Define struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// hasBirthday Method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++ //Increases age
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func (p Person) greet() string {
	return "Hello. My name is " + p.firstName + " " + p.lastName + ". I am " + strconv.Itoa(p.age)
}

func main() {
	//init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{firstName: "John", lastName: "Doe", city: "Boston", gender: "m", age: 29}
	//Alternative
	//person1 := Person{"Samantha", "Smith", "Boston", "f", 25}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Williams")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
