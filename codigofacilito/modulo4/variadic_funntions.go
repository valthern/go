package main

import "fmt"

// Variadic function
func promedio(calificaciones ...int) int {
	var promedio, suma int

	for _, calificacion := range calificaciones {
		suma = suma + calificacion
	}

	promedio = suma / len(calificaciones)

	return promedio
}

func main() {
	fmt.Println("Promedio:", promedio(10, 7, 8, 9, 9, 9, 7))
}
