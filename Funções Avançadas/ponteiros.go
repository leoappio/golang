package main

import "fmt"

func invertNumberByReference(number *int) {
	*number = *number * -1
}

func main() {
	number := 10
	invertNumberByReference(&number)
	fmt.Println(number)
}
