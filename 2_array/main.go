package main

import "fmt"

func main() {
	// var fruitArr [2]string

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Banana"

	fruitArr := []string{"Apple", "Banana", "Orange", "Grapes"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[2])   //Orange
	fmt.Println(len(fruitArr)) // returns length
	fmt.Println(fruitArr[1:3]) //banana orange
	fmt.Println(fruitArr[1:2]) // banana

}
