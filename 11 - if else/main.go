package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	number := 10

	if number > 15 {
		fmt.Println("O número é maior que 15")
	} else {
		fmt.Println("O número é menor ou igual a 15")
	}

	if otherNumber := 30; otherNumber > 0 {
		fmt.Printf("O numero é maior que 0: %d\n", otherNumber)
	} else if otherNumber < 0 {
		fmt.Println("O número é menor que 0")
	} else {
		fmt.Println("O número é igual a 0")
	}
}
