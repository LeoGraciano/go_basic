package main

import (
	"fmt"

	"github.com/badoux/checkmail"

	"modulo/auxiliar"
)

// Escrever registra uma mensagem na tela
func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Writer()
	err := checkmail.ValidateFormat("leonardoferreiragraciano@gmail.com")
	if err != nil {
		fmt.Println("Email inválido")
	} else {
		fmt.Println("Email válido")
	}
}
