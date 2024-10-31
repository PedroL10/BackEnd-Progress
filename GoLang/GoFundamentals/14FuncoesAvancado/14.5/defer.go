package main

import "fmt"

func funcao1() {
	fmt.Println("Executando funcao 1")
}

func funcao2() {
	fmt.Println("Executando funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado sera retornado")
	fmt.Println("entrando na funçao para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	} else {
		return false
	}

}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 7))
}
