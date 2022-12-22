package main

import (
	"fmt"
)

func main() {

	var sub1, sub2, sub3, sub4, sub5 int
	fmt.Println("Enter marks of the five subjects:")
	// fmt.Scanln(&sub1, &sub2, &sub3, &sub4, &sub5)    //Takes multiple input

	// fmt.Scanln("%d",&sub1)
	// fmt.Scan("%d", &sub2)
	// fmt.Scan("%d", &sub3)
	// fmt.Scan("%d", &sub4)
	// fmt.Scan("%d", &sub5)

	avg := (sub1 + sub2 + sub3 + sub4 + sub5) / 5
	if avg >= 90 {
		print("Grade: A")
	} else if avg >= 80 && avg < 90 {
		print("Grade: B")
	} else if avg >= 70 && avg < 80 {
		print("Grade: C")
	} else if avg >= 60 && avg < 70 {
		print("Grade: D")
	} else {
		print("Grade: F")
	}

	//Only takes a single input
	// var myString string
	// fmt.Scanln(&myString)
	// fmt.Println(myString)

	//Outputs the full name
	// var name string = "john doe"
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the rating: ")
	// myRating, _ := reader.ReadString('\n')
	// myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	// fmt.Println(myNumRating + 2)
}
