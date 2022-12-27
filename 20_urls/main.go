package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to URL")
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("Type is %T \n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	partsofurl := &url.URL{ //important way of handling URLs
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsofurl.String()
	fmt.Println(anotherUrl)

}
