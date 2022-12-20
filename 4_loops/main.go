package main

import "fmt"

func main() {

	//For Loop
	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			fmt.Printf("\n %d is divisible by 3", i)
		} else if i%5 == 0 {
			fmt.Printf("\n %d is divisible by 5", i)
		} else {
			fmt.Printf("\n %d is not divisible by 3 or 5", i)
		}
	}
}
