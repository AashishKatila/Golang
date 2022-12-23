package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 5
}

func main() {

	var ptr *int
	fmt.Println("Value of ptr = ", ptr)

	a := 5
	ptr = &a
	fmt.Println("Value of ptr = ", ptr)
	fmt.Println("Value of ptr = ", *ptr)

	i := 3
	fmt.Println("Value of i = ", i)

	zeroval(i)
	fmt.Println("Value of i = ", i)

	zeroptr(&i)
	fmt.Println("Value of i = ", i)

}
