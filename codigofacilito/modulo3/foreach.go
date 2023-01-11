package main

import (
	"fmt"
)

func main() {
	animales := [...]string{"Perro", "Gato", "Pez", "Vaca", "Cerdo"}

	/*
		for indice := 0; indice < len(animales); indice++{
			elemento := animales[indice]
			fmt.Println(elemento)
		}
	*/
	
	// foreach con dos variables
	for indice, elemento := range animales {
		fmt.Println("El indice es:", indice, ", el valor es:", elemento)
	}

	fmt.Println()

	// foreach con una variable
	for _, elemento := range animales {
		fmt.Println("El valor es:", elemento)
	}
}
