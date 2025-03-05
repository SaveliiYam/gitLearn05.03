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
	Name     string
	LastName string
	Age      int
}

func (u *User) PrintInfo() {
	fmt.Printf("Name: %s %s, Age: %d\n", u.Name, u.LastName, u.Age)
}
func (u User) Greet() {
	fmt.Printf("Hello, my name is %s %s! I'm %d years old.\n", u.Name, u.LastName, u.Age)
}
