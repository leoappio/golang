package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i")
	}

	for j := 0; j < 10; j++ {
		fmt.Println("j está sendo incrementado")
	}

	nomes := [3]string{"teste", "leonardo", "appio"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
}
