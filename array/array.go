package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitLIst [4]string

	fruitLIst[0] = "Apple"
	fruitLIst[1] = "Mango"
	fruitLIst[3] = "Peach"

	fmt.Println("length of fruitLIst is ", len(fruitLIst))

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("List of those vegetables are ", vegList)

}
