package main

import (
	"fmt"
)

func main() {
	fmt.Println("Funcoes nomeadas")
	soma, subtracao := calculosMatematicos(10, 7)
	fmt.Println(soma, " | ", subtracao)
}

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
