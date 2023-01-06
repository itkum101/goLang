package main

import (
	"fmt"
	"sort"
)

// Super easy and super important
func main() {
	fmt.Println("Welcome to slices")
	fmt.Println("Slices are amazing")

	// syntax of a slices
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	// We can add many values as we like and memory is expanded automatically

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// colon is for slicing

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 235
	// append will reaccomade the value

	highScores = append(highScores, 556, 666, 432)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	highScores = append(highScores[0:2], highScores[3:]...)
	fmt.Println(highScores)

}
