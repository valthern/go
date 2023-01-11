package main

import "fmt"

func main() {
	// se puede omitir el valor del tamaño inicial del arreglo con los tres puntos
	numeros1 := [5]int{100, 200, 300, 400, 500}
	numeros2 := [...]int{100, 200, 300, 400, 500}
	// Se pueden asignar con los índices el valor que va a llevar
	monedas := [...]string{0: "Peso mexicano", 1: "Dólar", 2: "Euro", 5: "Dólar canadiense"}
	subMonedas := monedas[0:3]

	fmt.Println(numeros1)
	fmt.Println(numeros2)
	fmt.Println(monedas[0])
	fmt.Println(monedas[1])
	fmt.Println(monedas[2])
	fmt.Println(monedas[3])
	fmt.Println(monedas[4])
	fmt.Println(monedas[5])
	fmt.Println(subMonedas)
}
