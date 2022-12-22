package main

import (
	"fmt"
)

//Count set bits

func main() {
	var num int
	rem := 0
	quo := 0
	final := 0
	fmt.Printf("Enter Number to convert to binary and count set bits: \n")
	fmt.Scanf("%d", &num)
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
	fmt.Printf("%d \n", final)

	count := 0
	for final != 0 {
		remainder := final % 10
		if remainder == 1 {
			count++
		}
		final = final / 10
	}
	fmt.Printf("Set Bits = %d", count)
}
