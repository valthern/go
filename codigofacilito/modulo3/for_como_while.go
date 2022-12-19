package main

import "fmt"

func main() {
	numero := 12345
	contador := 0

	// Funciona como "while"
	for numero > 0 {
		numero = numero / 10
		contador++
	}

	fmt.Println("La cantidad de dígitos es:", contador)
}