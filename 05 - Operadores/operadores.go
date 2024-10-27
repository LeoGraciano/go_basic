package main

func main() {
	somar := 1 + 2
	subtrair := 5 - 3
	multiplicar := 4 * 6
	dividir := 12 / 3
	restoDaDivisão := 7 % 2
	println(somar, subtrair, multiplicar, dividir, restoDaDivisão)

	var num1 int16 = 10
	var num2 int16 = 20
	somaComCast := num1 + num2
	println(somaComCast)

	// Operadores Relacionais
	var num3 int = 10
	var num4 int = 20
	maior := num3 > num4
	menor := num3 < num4
	igual := num3 == num4
	diferente := num3 != num4
	println(maior, menor, igual, diferente)

	// Operadores Lógicos
	println("---------------------")
	var verdadeiro bool = true
	var falso bool = false
	e := verdadeiro && falso
	ou := verdadeiro || falso
	nao := !verdadeiro
	println(e, ou, nao)

	// Operadores Unários
	println("---------------------")
	var numero int = 10
	numero++
	incremento := numero
	numero--
	decremento := numero
	println(incremento, decremento)

	// Operadores Aritméticos
	println("---------------------")
	somaAritmetica := numero + incremento
	subtracaoAritmetica := numero - decremento
	multiplicacaoAritmetica := numero * incremento
	divisaoAritmetica := numero / decremento
	restoDaDivisaoAritmetica := numero % decremento
	println(somaAritmetica, subtracaoAritmetica, multiplicacaoAritmetica, divisaoAritmetica, restoDaDivisaoAritmetica)

	// Operadores de Associação
	println("---------------------")
	var numero2 int = 10
	numero2 += incremento
	somaComOperadorDeAtribuição := numero2
	numero2 -= decremento
	subtraçãoComOperadorDeAtribuição := numero2
	numero2 *= incremento
	multiplicaçãoComOperadorDeAtribuição := numero2
	numero2 /= decremento
	divisãoComOperadorDeAtribuição := numero2
	restoDaDivisãoComOperadorDeAtribuição := numero2 % decremento
	println(somaComOperadorDeAtribuição, subtraçãoComOperadorDeAtribuição, multiplicaçãoComOperadorDeAtribuição, divisãoComOperadorDeAtribuição, restoDaDivisãoComOperadorDeAtribuição)

}
