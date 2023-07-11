package main

import "fmt"

/* Con defer se ejecutan en orden de pila, es decir 
	ingresan dentro de una pila y la última que se 
	añadió es la primera en salir (LIFO)
*/

func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	a:=5
	defer fmt.Println("Defer:",a)

	a=10
	fmt.Println(a)
}