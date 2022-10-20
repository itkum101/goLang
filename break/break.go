package main

import "fmt"

func main() {

	fmt.Println("Welcome to breaks on golang")

	days := []string{"Sunday", "Tuesday", "Wednesday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, value := range days {
		fmt.Printf("index is %v and value is %v\n", index, value)

	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 5 {
			break

		}
		fmt.Println("Value is ", rogueValue)
		rogueValue++
	}

}
