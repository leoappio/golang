package main

import "fmt"

func main() {
	var variavelString string = "teste string"
	variavelImplicita := "teste var implicita"
	fmt.Println(variavelString)
	fmt.Println(variavelImplicita)

	var (
		variavel3 string = "teste"
		variavel4 string = "teste 2"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "teste", "teste"
	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
}
