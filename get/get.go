package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// PerformPostJsonRequest()
	PerformGetRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
	{
		"coursename":"Let's go with go lang",
		"price":0,
		"platform":"leetCodeOnline.in"
		
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic("Error occured")
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("STATUS CODE: ", response.StatusCode)
	fmt.Println("Content length is ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	// Another method
	var responseString strings.Builder
	// It gives out the no of byte in our word
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

}
