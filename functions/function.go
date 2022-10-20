package main

import "fmt"

func main() {
	fmt.Println("Welcome to function ")

	greeter()
	result, _ := proAdder(3, 5)

	fmt.Println("Result is : ", result)

}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Message embeeded with total"
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("WElcome man")
}
