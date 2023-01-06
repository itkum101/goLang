package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("WEB REQUEST")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Typef of response is %T \n", response.Body)
	fmt.Println(response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)

}
