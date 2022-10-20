package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("\n\n")
	fmt.Println("Welcome to file tutorial !!! !!!")
	content := "This needs to go in a file "
	fmt.Println(content)
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./mylcogofile.txt")

	fmt.Printf("\n\n")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
