package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// type Message struct {
// 	Name string
// 	Body string
// 	Time int64
// }

func main() {
	fmt.Println("Creating and Encoding JSON Data")
	// m := Message{"John", "Hello", 1294706395881547000}

	// b, err := json.Marshal(m)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(b)
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React-js", 266, "hello.com", "abc123", []string{"web-dev", "js"}},
		{"MERN", 166, "hello.mern.com", "bcd123", []string{"full-stack", "js"}},
		{"Angular-bootcamp", 266, "example.com", "bcd456", nil},
	}
	//package this data to json data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "React-js",   
		"Price": 266,
		"website": "hello.com",     
		"tags": ["web-dev","js"]
	}
		`)
	// fmt.Println(jsonDataFromWeb)
	var goCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON valid")
		json.Unmarshal(jsonDataFromWeb, &goCourse)
		fmt.Printf("%#v \n", goCourse)
	} else {
		fmt.Println("JSON not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v adn value is %v and Type: %T \n", k, v, v)
	}

}
