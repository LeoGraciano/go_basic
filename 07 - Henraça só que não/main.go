package main

import "fmt"

type Person struct {
	name string
	age  int8
}

type Study struct {
	Person
	course string
}

func main() {
	fmt.Println("Hello, World!")

	var p Person
	p.name = "John Doe"
	p.age = 30
	fmt.Println(p)

	var s Study
	s.name = "Jane Doe"
	s.age = 25
	s.course = "Computer Science"
	fmt.Println(s)
}
