package main

import "fmt"

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execucao")
}

func recuperarExecucao() {
	if r := recover(); r != nil {

		fmt.Println("execucao recuperada")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A media é exatamente 6")

}
