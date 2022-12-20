package main

import (
	"fmt"
)

//Decimal into Binary using while loop

func main() {
	num := 23
	rem := 0
	quo := 0
	final := 0
	for num != 0 {
		// fmt.Println(num)
		// num--
		rem = rem*10 + num%2
		quo = num / 2
		num = quo
	}

	fmt.Printf("%d \n", rem)

	for rem != 0 {
		final = final*10 + rem%2
		rem = rem / 10
	}
	fmt.Printf("%d", final)
}
