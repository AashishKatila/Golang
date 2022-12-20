package main

import "fmt"

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello")
	fmt.Println(getSum(1, 9))
}
