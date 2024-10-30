package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("É maior que zero")
	} else if numero < -10 {
		fmt.Println("É menor do que -10")
	} else {
		fmt.Println("Entre 0 e 10")

	}

}
