package main

import "fmt"

type User struct {
	name string
	age  uint8
}

func main() {
	var u User
	u.name = "John Doe"
	u.age = 50
	fmt.Println(u)

	user1 := User{name: "John Doe", age: 30}
	user2 := User{"Jane Doe", 25}
	fmt.Println(user1, user2)
}
