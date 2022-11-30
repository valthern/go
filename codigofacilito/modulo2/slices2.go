package main

import "fmt"

/*
Slices:  Se componen de los siguientes elementos
1.- Un puntero
2.- Una longitud
3.- Una capacidad
*/

func main() {
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio",
		"Agosto", "Septiembre"} //,"Octubre","Noviembre","Diciembre"}

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("La longitud es: %v, La capacidad es: %v %p\n", longitud, capacidad, meses)

	// Si la estructura se encuentra en su capacidad máxima
	meses = append(meses, "Octubre")
	longitud = len(meses)
	capacidad = cap(meses)
	
	fmt.Printf("La longitud es: %v, La capacidad es: %v %p\n", longitud, capacidad, meses)
	
	meses = append(meses, "Noviembre")
	meses = append(meses, "Diciembre")
	longitud = len(meses)
	capacidad = cap(meses)
	
	fmt.Printf("La longitud es: %v, La capacidad es: %v %p\n", longitud, capacidad, meses)
}
