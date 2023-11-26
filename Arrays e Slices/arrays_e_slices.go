package main

import "fmt"

func main() {
	var array1 [5]int

	array1[0] = 10

	fmt.Println(array1)

	array2 := [2]string{"teste", "teste2"}
	fmt.Println(array2)

	slice := []int{10, 12, 34, 45}

	//cria um slice novo com o elemento novo no final
	slice = append(slice, 36)

	fmt.Println(slice)

	//Arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade
}
