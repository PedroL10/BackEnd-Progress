package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Praça dos Imigrantes")
	fmt.Println(tipoEndereco)
}
