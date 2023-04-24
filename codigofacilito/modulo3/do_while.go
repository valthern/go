package main

import "fmt"

func main() {
	var numero = 10

	// Variable booleana, siempre "ok", y luego la condición asignada a la variable ok
	for ok := true; ok; ok = numero < 10 {
		fmt.Println(numero)
		numero++
	}
}
