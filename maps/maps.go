package main

import "fmt"

func main() {
	fmt.Println("WELCOME TO MAPS")

	languages := make(map[string]string)

	languages["JS"] = "JAVASCRIPT"
	languages["RB"] = "RUBY"
	languages["PY"] = "PYTHON"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS SHORTS FOR : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	for key, value := range languages {
		fmt.Printf("For key %v , value is %v \n", key, value)
	}
}
