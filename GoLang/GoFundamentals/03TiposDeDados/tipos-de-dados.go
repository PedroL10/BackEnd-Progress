package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int8 = 127
	var numero1 int16 = 32767
	var numero2 int32 = 2147483647
	var numero3 int64 = 922337203685477580

	numero4 := 999999999

	var numero5 uint = 44

	fmt.Println(numero)
	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)

	fmt.Println("=============================")

	var numeroReal float32 = 13232.33
	var numeroReal2 float64 = 412312212.33

	numeroReal3 := 1321.99

	fmt.Println(numeroReal)
	fmt.Println(numeroReal2)
	fmt.Println(numeroReal3)

	fmt.Println("=============================")

	var str string = "Texto teste"
	str2 := "Outro texto de teste"
	fmt.Println(str)
	fmt.Println(str2)

	char := 'A'

	fmt.Println(char)

	fmt.Println("=============================")

	booleano := false
	var booleano1 bool = true
	fmt.Println(booleano1)
	fmt.Println(booleano)

	fmt.Println("=============================")

	var erro error = errors.New("erro teste")
	fmt.Println(erro)

}
