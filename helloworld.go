// Package declaration
package main

import "fmt"

//fmt --> Formatting for input and output
type Course struct {
	CourseId string
}

func main() {

	c := Course{"Hello"}
	test(&c)
	fmt.Println(c.CourseId)

}
func test(c *Course) {
	// No need to dereferencing it .
	c.CourseId = "testing2"

}
