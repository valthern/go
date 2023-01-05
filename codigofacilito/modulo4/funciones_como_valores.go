package main

import "fmt"

func factorial(numero int) int {
	if numero == 1 || numero == 0 {
		return 1
	}

	return factorial(numero-1) * numero
}

func main() {
	var miFuncion = factorial

	resultado := miFuncion(3)

	fmt.Println(resultado)
}
