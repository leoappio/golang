package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	var resultado int = soma(10, 20)
	fmt.Println(resultado)
}
