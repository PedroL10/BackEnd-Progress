package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)
	c := circulo{10}
	escreverArea(c)

}