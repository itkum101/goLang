package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("worlds")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
