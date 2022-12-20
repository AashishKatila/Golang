package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var myString string
	// fmt.Scanln(&myString) //Only takes a single input
	// fmt.Println(myString)

	// var name string = "john doe"
	// var a, _ = fmt.Println(name)
	// fmt.Println(a) //outputs the full name

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating: ")
	myRating, _ := reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	fmt.Println(myNumRating + 2)
}

// 	var i int

// 	fmt.Print("Type a number: ")
// 	fmt.Scan(&i)
// 	fmt.Println("Your number is:", i)
