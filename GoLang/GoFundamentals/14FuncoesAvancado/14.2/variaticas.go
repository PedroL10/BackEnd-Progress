package main

import (
	"fmt"
)

func soma(numeros ...int) int {
	total := 0

	for i := 0; i < len(numeros); i++ {
		total = total + i
	}
	return total

	// for _, numero := range numeros {    MESMA FUNCIONALIDADE DO FOR ACIMA
	// 	total += numero
	// }
	// return total

}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
	fmt.Println(texto, numeros)
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 3, 4, 5, 3, 4, 5, 6, 4, 2, 3)
	totalDaSoma1 := soma(9, 9, 9, 9, 9)
	totalDaSoma2 := soma()
	fmt.Println(totalDaSoma)
	fmt.Println(totalDaSoma1)
	fmt.Println(totalDaSoma2)

	fmt.Println("=========================")

	escrever("Estudando Go", 1, 2, 3, 4, 5)

}
