package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}
func (this User) nombre_completo() string {
	return this.nombre + " " + this.apellido
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
	ramon.apellido="Fernández"
	fmt.Println(ramon.nombre_completo())
}
