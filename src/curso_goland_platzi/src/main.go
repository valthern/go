package main

import "fmt"

func main() {
	// Declaración de constantes. Hasta aquí llega la precisión del float64
	const pi float64 = 3.141592653589793
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)
}
