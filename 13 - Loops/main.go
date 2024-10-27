package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	i := 0

	for i < 10 {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Loop com range")
	arr := []int{10, 20, 30, 40, 50}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
		time.Sleep(time.Second)
	}
	fmt.Println("Loop com string")
	str := "Hello, World!"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c, String: %v\n", index, char, string(char))
		time.Sleep(time.Second)
	}

	user := map[string]string{
		"name":  "John Doe",
		"email": "johndoe@example.com",
	}
	for key, value := range user {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
		time.Sleep(time.Second)
	}

	// fmt.Println("Loop infinito com for")
	// for {
	// 	fmt.Println("Loop infinito")
	// 	time.Sleep(time.Second)
	// }

}
