package main

import (
	"fmt"
)

func calcMath(n1, n2 int) (soma int, substração int) {
	soma = n1 + n2
	substração = n1 - n2
	return
}

func somar(numbers ...int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(posição uint) uint {
	if posição <= 1 {
		return posição
	}
	return fibonacci(posição-2) + fibonacci(posição-1)
}

func closure() func() {
	text := "Dentro da função closure"
	function := func() {
		fmt.Println(text)
	}

	return function
}

func inverterSignal(n int) int {
	return n * -1
}

func inverterSignalPointer(n *int) {
	*n = *n * -1
}

func init() {
	fmt.Println("Função init() chamada")
	fmt.Println("----------------------")
}

func main() {
	fmt.Println("Retorno nomeado")
	soma, subtração := calcMath(10, 20)
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtração)
	fmt.Println("----------------------")
	fmt.Println("Funções Variáticas")
	fmt.Println("Soma:", somar(10, 20, 30, 40, 50))
	fmt.Println("----------------------")
	fmt.Println("Funções anônimas")
	func(text string) {
		fmt.Println(text)
	}("passando função anônima")
	fmt.Println("----------------------")
	fmt.Println("Funções Recursivas")
	fmt.Println("Fatorial de 5:", factorial(5))
	fmt.Println("Fatorial de 10:", factorial(10))

	posição := uint(10)
	fmt.Printf("Fibonacci de %dº posição: %d\n", posição, fibonacci(posição))
	fmt.Println("----------------------")
	fmt.Println("Função Defer")
	defer fmt.Println("Executado após a função terminar")
	fmt.Println("----------------------")
	fmt.Println("Função Closure")
	text := "Dentro da função main"
	println(text)
	newFunction := closure()
	newFunction()
	fmt.Println("----------------------")
	fmt.Println("Função com Ponteiro")
	number := 20
	invert := inverterSignal(number)
	fmt.Println("Número invertido:", invert)
	fmt.Println("Número:", number)
	inverterSignalPointer(&number)
	fmt.Println("Número invertido com Ponteiro:", number)

	fmt.Println("----------------------")
	fmt.Println("Função init")

	fmt.Println("----------------------")
	fmt.Println("Panic e Recovery")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado do panic:", r)
		}
	}()
	panic("Gerando panic")
}
