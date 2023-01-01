package main

import "fmt"

// La primera forma es tradicional si hay dos parámetros de diferente tipo
// La forma segunda se usa si los dos parámetros son del mismo tipo
// func suma(numero1 int,numero2 int) int {
func suma0(numero1, numero2 int) (int, string) {
	return numero1 + numero2, "Un mensaje desde la función suma"
}

func main() {
	//fmt.Println(suma(10, 20))

	resultado, mensaje := suma0(10, 20)

	fmt.Println(resultado, mensaje)
}
