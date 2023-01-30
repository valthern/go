package main

import "fmt"

func factorial(numero int) int {
	if numero == 1 || numero == 0 {
		return 1
	}

	return factorial(numero-1) * numero
}

type customFunction func(n int) int

func main() {
	//var miFuncion = factorial
	//var miFuncion func(n int) int
	var miFuncion customFunction //nil al declarar la variable

	if miFuncion == nil {
		miFuncion = factorial
	}

	resultado := miFuncion(3)

	fmt.Println(resultado)
}
