package main

import "fmt"

func main() {
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio",
		"Agosto", "Septiembre"} //,"Octubre","Noviembre","Diciembre"}

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("La longitud es: %v, La capacidad es: %v\n", longitud, capacidad)
	
	meses = append(meses, "Octubre")
	longitud = len(meses)
	capacidad = cap(meses)
	
	fmt.Printf("La longitud es: %v, La capacidad es: %v\n", longitud, capacidad)

}
