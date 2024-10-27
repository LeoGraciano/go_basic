package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":  "John Doe",
		"email": "johndoe@example.com",
	}
	fmt.Println(user["name"])

	users := map[string]map[string]string{
		"user1": {
			"name":  "John Doe",
			"email": "johndoe@example.com",
		},
		"user2": {
			"name":  "Jane Doe",
			"email": "janedoe@example.com",
		},
	}
	fmt.Println(users)
	delete(users, "user1")
	fmt.Println(users)
}
