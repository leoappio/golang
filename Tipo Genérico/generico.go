package main

import "fmt"

func print(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	print("teste")
	print(2)
}
