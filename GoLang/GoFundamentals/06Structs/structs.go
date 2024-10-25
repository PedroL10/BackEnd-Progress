package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {

	var u usuario
	u.idade = 31
	u.nome = "Pedro"
	u.endereco = endereco{"Rua teste", 99}

	enderecoExemplo := endereco{"Rua das flores", 55}

	fmt.Println(u)

	usuario2 := usuario{"Teste", 22, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Teste sem idade"}
	fmt.Println(usuario3)

}
