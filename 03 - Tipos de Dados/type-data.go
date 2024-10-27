package main

import "fmt"

func main() {
	var number1 int = 1000000000000000000
	fmt.Printf("O número é: %d\n", number1)
	number2 := 999999999999999999
	fmt.Printf("O número é: %d\n", number2)

	var number3 uint32 = 1000000000
	fmt.Printf("O número é: %d\n", number3)

	//alias
	//INT32 == RUNE
	var number4 rune = 123456
	fmt.Printf("O número é: %d\n", number4)

	//byte = uint8
	var number5 byte = 123
	fmt.Printf("O número é: %d\n", number5)

	//float32, float64
	var number6 float32 = 123.456
	fmt.Printf("O número é: %.2f\n", number6)

	var number7 float64 = 123.456
	fmt.Printf("O número é: %.2f\n", number7)

	//complex64, complex128
	var number8 complex64 = 123.456 + 78.901i
	fmt.Printf("O número é: %.2f + %.2fi\n", real(number8), imag(number8))

	var number9 complex128 = 123.456 + 78.901i
	fmt.Printf("O número é: %.2f + %.2fi\n", real(number9), imag(number9))

	//String
	var str1 string = "Hello, World!"
	fmt.Printf("O string é: %s\n", str1)

	char1 := 'H'
	fmt.Println("O caractere é: ", char1)

	//boolean
	var bool1 bool = true
	fmt.Printf("O boolean é: %t\n", bool1)
	bool2 := false
	fmt.Printf("O boolean é: %t\n", bool2)

	//slice
	var slice1 []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("O slice é: %v\n", slice1)

	//map
	var map1 map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("O map é: %v\n", map1)

	//struct
	type Person struct {
		Name string
		Age  int
	}
	var person1 Person = Person{"John Doe", 30}
	fmt.Println("A pessoa é", person1)

	//error
	var err error = fmt.Errorf("Ocorreu um erro")
	fmt.Println("O erro é:", err)
}
