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
	fmt.Println(dog, cat, face)

	// a es de tipo rune, el cual es int32 y tambien caracter unicode
	var a rune = 97
	fmt.Printf("Tipo: %T, Valor: %c\n", a, a)

	var x uint8 = 100
	var y uint16 = 23000
	// z es casteada
	z := uint16(x) + y
	fmt.Printf("Tipo: %T, Valor: %v\n",z,z)
	// Pi es una constante que vale 3.14159265358979...
	//const Pi = 3.14159265358979
}
