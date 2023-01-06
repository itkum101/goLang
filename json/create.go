package main

import (
	"encoding/json"
	"fmt"
)

// Creating structure

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"_"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("Json")
	DecodeJson()
	//EncodeJson()

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		
		{
			"Name": "ReactJS ",
			"Price": 299,
			"Platform": "LCO",
			"Password": "abc123",
			"Tags": ["web-dve","Js"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was a valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID MAN")
	}

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("\n\n")
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}

}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS ", 299, "LCO", "abc123", []string{"web-dve", "Js"}},
		{"MERN ", 299, "LCO", "abc123", []string{"web-dve", "Js"}},
		{"Angular ", 299, "LCO", "bcd123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
