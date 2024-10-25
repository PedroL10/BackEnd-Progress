package main

import "fmt"

func main() {
	var variavel1 string = "variavel"
	var variavel2 int = 99
	var variavel3 float32 = 3.92929

	variavel4 := "teste"
	variavel5 := 10
	variavel5 = 22
	variavel6 := true

	var (
		variavel7 string = "abcd"
		variavel8 string = "lalala"
	)

	fmt.Println("variavel 1:", variavel1)
	fmt.Println("variavel 2:", variavel2)
	fmt.Println("variavel 3:", variavel3)
	fmt.Println("variavel 4:", variavel4)
	fmt.Println("variavel 5:", variavel5)
	fmt.Println("variavel 6:", variavel6)
	fmt.Println("variavel 7:", variavel7)
	fmt.Println("variavel 8:", variavel8)

	const constante1 string = "String constante"
	fmt.Println("constante 1:", constante1)
}
