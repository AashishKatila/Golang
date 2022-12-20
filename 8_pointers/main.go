package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Println(a)
	fmt.Printf("%T\n", b)
	fmt.Println(b)
	fmt.Println(*b)
	*b = 23
	fmt.Println(*b)
	fmt.Println(*&a)
}
