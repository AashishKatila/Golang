package main

import "fmt"

func main() {
	// var fruits [2]string //Array

	// fruits[0] = "Apple"
	// fruits[1] = "Banana"
	// fmt.Println(fruits)

	// fruitArr := []string{"Apple", "Banana", "Orange", "Grapes"} //Slice

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[2])   //Orange
	// fmt.Println(len(fruitArr)) // returns length
	// fmt.Println(fruitArr[1:3]) //banana orange
	// fmt.Println(fruitArr[1:2]) // banana
	// fruitArr = append(fruitArr, "Litchi", "Mango")
	// fmt.Println(fruitArr)

	nums := make([]int, 3, 3)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 9
	fmt.Println("nums = ", nums)
	nums = append(nums, 3, 4, 5)
	fmt.Println(nums)

	var index int = 2
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println(nums)
}
