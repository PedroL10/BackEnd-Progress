package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("pep@gmail.com")
	erro1 := checkmail.ValidateFormat("pep")
	fmt.Print(erro)
	fmt.Print(erro1)
}
