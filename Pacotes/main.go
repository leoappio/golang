package main

import (
	"fmt"
	auxiliar "modulo/Auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Print do modulo main")

	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("leonardo.appio@gmail.com")

	fmt.Println(erro)
}
