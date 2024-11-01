package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//CONCORRENCIA != PARALELISMO

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1 go routine
	}()

	go func() {
		escrever("GO")
		waitGroup.Done() // -1 go routine
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
