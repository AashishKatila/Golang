package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://aashishkatila.com.np"

func main() {
	fmt.Println("Learning Web Request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Type =  ", response)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(data)
	fmt.Println(content)
}
