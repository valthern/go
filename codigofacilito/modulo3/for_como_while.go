package main

import "fmt"

func main() {
	numero := 12345
	contador := 0

	// Funciona como "while". Sólo contiene una parte que es la condicional
	for numero > 0 {
		numero = numero / 10
		contador++
	}

	fmt.Println("La cantidad de dígitos es:", contador)
}