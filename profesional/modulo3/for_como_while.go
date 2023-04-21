package main

import "fmt"

func main() {
	numero := 12345
	contador := 0

	// Funciona como "while". Sólo contiene la condición. Se usa como un "while" tal cual.
	for numero > 0 {
		numero = numero / 10
		contador++
	}

	fmt.Println("La cantidad de dígitos es:", contador)
}