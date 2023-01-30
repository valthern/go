package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var curso string = "Hola desde el curso profesional de Go!" // 1
	// var curso = "Curso profesional de Go!" // 2
	curso := "Curso profesional de Go!" // 3

	longitud := len(curso) // int

	fmt.Println(longitud)

	primerCaracter := curso[0] // char -> uint8
	ultimoCaracter := curso[len(curso)-1]

	fmt.Println(primerCaracter)
	fmt.Println(reflect.TypeOf(primerCaracter))

	fmt.Printf("%c\n", primerCaracter)
	fmt.Printf("%c\n", ultimoCaracter)
}
