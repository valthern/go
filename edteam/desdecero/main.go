// package main es el paquete principal
// que me permite ejecutar mi aplicación
package main

import "fmt"

// func main() es la funcion principal
func main() {
	// dog es la primera variable de prueba
	// cat es la segunda variable de prueba
	dog, cat := "🐕", "😺"
	// face es la tercera variable de prueba
	cat, face := "gatico", "🥴"

	// Pi es una constante que vale 3.14159265358979...
	const Pi = 3.14159265358979

	fmt.Println(dog, cat, face)
}
