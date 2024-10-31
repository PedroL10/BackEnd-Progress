package main

import "fmt"

var n int

func main() {
	fmt.Println("Main sendo executada")
	fmt.Print("valor de n ", n)
}

func init() {
	n = 10
	fmt.Println("Executando antes da main")
}
