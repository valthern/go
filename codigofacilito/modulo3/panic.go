package main

import "fmt"

func main() {
	// funcion panic

	var dividendo, divisor int

	fmt.Print("Ingresa un valor para el dividendo: ")
	fmt.Scanf("%d", &dividendo)

	fmt.Print("Ingresa un valor para el divisor: ")
	fmt.Scanf("%d", &divisor)

	if divisor == 0{
		panic("No es posible dividir sobre cero")
	}

	resultado := dividendo / divisor

	fmt.Println("Resultado: ", resultado)
}
