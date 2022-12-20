package main

import "fmt"

func main() {
	//If Else
	color := "blue"

	if color == "red" {
		fmt.Printf("Color is Red \n")
	} else if color == "blue" {
		fmt.Printf("Color is Blue \n")
	} else {
		fmt.Printf("Color is neither Red nor Blue \n")
	}

	//Switch Case
	OS := "macos"
	switch OS {
	case "windows":
		fmt.Printf("User uses Windows \n")
	case "macos":
		fmt.Printf("User uses MacOS \n")
	default:
		fmt.Printf("User uses Linux \n")
	}
}
