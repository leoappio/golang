package main

import "fmt"

type retangulo struct {
	altura  float32
	largura float32
}

type forma interface {
	area() float32
}

func (r retangulo) area() float32 {
	return r.altura * r.largura
}

func imprimirArea(f forma) {
	fmt.Println(f.area())
}

func main() {
	retangulo1 := retangulo{float32(10.3), float32(45.6)}
	imprimirArea(retangulo1)
}
