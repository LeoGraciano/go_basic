package main

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func calcMath(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub

}

func main() {
	result := somar(10, 20)
	println(result)

	resultSoma, resultSub := calcMath(10, 20)
	println(resultSoma, resultSub)
}
