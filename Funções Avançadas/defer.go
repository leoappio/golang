package main

import "fmt"

func funcao1() {
	fmt.Println("executando a funcao1")
}

func funcao2() {
	fmt.Println("executando a funcao2")
}

func main() {
	//DEFER == ADIAR
	//irá executar no ultimo momento possível
	defer funcao1()
	funcao2()

}
