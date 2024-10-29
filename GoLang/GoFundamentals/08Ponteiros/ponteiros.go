package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, " | ", variavel2)

	variavel1++
	variavel2--

	fmt.Println(variavel1, " | ", variavel2)

	//PONTEIRO é uma referencia de memoria - armazena o endereco de memoria
	var variavel3 int
	var ponteiro *int

	variavel3 = 150
	ponteiro = &variavel3

	fmt.Println(variavel3, " | ", ponteiro, " | ", *ponteiro) // **desreferenciacao
}
