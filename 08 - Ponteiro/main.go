package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ponteiros")
	var var1 int = 10
	var var2 = var1
	fmt.Println("Valor de var1:", var1)
	fmt.Println("Valor de var2:", var2)

	var1++
	fmt.Println("Valor de var1:", var1)
	fmt.Println("Valor de var2:", var2)

	// Ponteiros
	var ptr1 *int = &var1
	var ptr2 = ptr1
	fmt.Println("Valor de ptr1:", ptr1)
	fmt.Println("Valor de ptr2:", ptr2)

	*ptr1++
	fmt.Println("Valor de ptr1:", ptr1)
	fmt.Println("Valor de ptr2:", ptr2)
	fmt.Println("Valor de var1:", var1)

	// Referências
	var ref1 = &var1
	var ref2 = ref1
	fmt.Println("Valor de ref1:", ref1)
	fmt.Println("Valor de ref2:", ref2)

	*ref1 = 20
	fmt.Println("Valor de ref1:", ref1)
	fmt.Println("Valor de ref2:", ref2)
	fmt.Println("Valor de var1:", var1)
}
