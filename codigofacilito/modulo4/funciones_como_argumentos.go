package main

import "fmt"

type Operacion func(balance, cantidad int) int

func suma(numero1, numero2 int) int {
	return numero1 + numero2
}

func resta(numero1, numero2 int) int {
	return numero1 - numero2
}

func ejecutarOperacion(funcion Operacion, balance, cantidad int) {
	fmt.Println("Antes de la operación")

	resultado := funcion(balance, cantidad)
	fmt.Println("El resultado es:", resultado)
	fmt.Println("Después de la operación")
}

func main() {
	ejecutarOperacion(suma, 10,20)
	ejecutarOperacion(resta, 10,20)
}
