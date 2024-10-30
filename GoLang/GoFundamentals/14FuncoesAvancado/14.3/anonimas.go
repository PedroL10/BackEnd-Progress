package main

import "fmt"

func main() {

	func() {
		fmt.Println("Funcao anonima 1")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Funcao anonima 2")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Funcao anonima 3")

	fmt.Println(retorno)
}
