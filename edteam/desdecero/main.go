// package main es el paquete principal
// que me permite ejecutar mi aplicación
package main

import "fmt"

func main() {
	// dog, cat y face son variables de prueba
	dog, cat := "🐕", "😺"
	cat, face := "gatico", "🥴"

	// Pi es una constante que vale 3.14159265358979...
	const Pi = 3.14159265358979

	fmt.Println(dog, cat, face)
}
