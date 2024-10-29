package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Luis",
		"telefone":  "2432423",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Loreta",
			"ultimo":   "Andrade",
			"telefone": "343335545",
		},
		"curso": {
			"nome":   "Psicologia",
			"campus": "Puc",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
