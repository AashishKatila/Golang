package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Writing contents in file")

	content := "This is a new text added."

	file, err := os.Create("./newFile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length : ", length)
	defer file.Close()
	readFile("./newFile.txt")
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Datas in the file = ", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
