package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	for {
		mensagem, aberto := <-canal
		fmt.Println(aberto)
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
