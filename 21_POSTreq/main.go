package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	values := map[string]string{"name": "John Doe", "occupation": "gardener"}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])

	// performPostJSONReq()
}

func performPostJSONReq() {
	const myUrl = "http://www.localhost:8000/post"
	requestBody := strings.NewReader(`
		{
			"coursename":"Golang",
			"price": 100,
			"platform" : "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
