package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var user usuario
	user.idade = 21
	user.nome = "Leonardo Lima Appio"
	fmt.Println(user)

	usuario2 := usuario{"teste", 21}
	fmt.Println(usuario2)

}
