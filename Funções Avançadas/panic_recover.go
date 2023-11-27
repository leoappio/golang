package main

import "fmt"

func recoveryExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func throwPanic() {
	defer recoveryExecution()
	panic("teste de panic")
}

func main() {
	throwPanic()
}
