package main

import "fmt"

func main() {
	qwerty()
	aPlusB()
}

func qwerty() {
	fmt.Println("Hello, world!")
}

func aPlusB() {
	fmt.Println(3 + 3)
}

func asd() {
	fmt.Println(4 + 4)
}

type User struct {
	Name string
	Age  int
}
