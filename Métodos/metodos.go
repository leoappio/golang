package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
}

func (u usuario) salvar() {
	fmt.Println("exemplo de metodo que pode ser chamado atráves do usuario")
}

func (u *usuario) alterarNome(novoNome string) {
	u.nome = novoNome
}

func main() {
	usuario1 := usuario{"leonardo", "appio"}
	usuario1.salvar()

	usuario1.alterarNome("leonardo alterado")
}
