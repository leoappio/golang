package main

import "fmt"

func main() {

	//declara e já executa a função anonima
	func() {
		fmt.Println("teste de função anonima")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("passando como parametro")

	fmt.Println(retorno)
}
