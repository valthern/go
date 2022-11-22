package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func main() {
	// var omar User
	fmt.Println(User{nombre: "Omar", apellido: "Martínez", edad: 44})

	omar := User{nombre: "Omar", apellido: "Martínez"}
	fmt.Println(omar)

	uriel := User{33, "Uriel", "Carrillo"}
	fmt.Println(uriel)

	ramon := new(User)
	ramon.nombre = "Ramón"
	fmt.Println(ramon.nombre)
}
