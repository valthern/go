package main

import "fmt"

func mostrarVariable(objeto interface{})  {
	fmt.Printf("El valor de la variable es: %v\n",objeto)
}

func suma(n1, n2 int) int {
	return n1 + n2
}

type User struct {
	Nombre string
}

func main()  {
	usuario := User {Nombre: "Alvaro"}

	mostrarVariable(10)
	mostrarVariable(true)
	mostrarVariable("cadena")
	mostrarVariable(suma)
	mostrarVariable(usuario)
}