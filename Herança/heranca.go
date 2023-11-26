package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso string
}

func main() {
	var estudante estudante

	estudante.nome = "teste"
	estudante.curso = "curso"

	fmt.Println(estudante)
}
