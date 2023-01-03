package main

import "fmt"

func modificarValor(parametro *string) {
	*parametro = "¡¡¡Cambio de valor!!!"
}

func main() {
	var curso = "Curso profesional de Go!"

	fmt.Println("Antes de la función:", curso)

	modificarValor(&curso)

	fmt.Println("Después de la función:", curso)
}
