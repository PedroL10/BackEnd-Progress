package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (user usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", user.nome)
}

func (user usuario) verificaMaioridade() bool {
	return user.idade >= 18
}

func (user *usuario) fazerAniversario() {
	user.idade++
}

func main() {
	usuario1 := usuario{"Usuario teste", 10}

	usuario2 := usuario{"Teste", 30}
	usuario1.salvar()
	usuario2.salvar()
	fmt.Println(usuario1.verificaMaioridade())
	fmt.Println(usuario2.verificaMaioridade())

	fmt.Println(usuario2.idade)
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
