package main

import "fmt"

func main() {
	numero := 10

	if numero == 10 {
		fmt.Println("igual a 10")
	} else {
		fmt.Println("nao é maior que 10")
	}

	//if init - cria uma variavel temporaria limitada ao escopo do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que zero")
	}

	switch numero {
	case 10:
		fmt.Println("é 10")
	default:
		fmt.Println("não é 10")
	}

	switch {
	case numero == 10:
		fmt.Println("é 10")
	}
}
