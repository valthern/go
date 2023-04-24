package main

import "fmt"

/* Enum */
type UserType int

const (
	Maestro    UserType = 1
	Estudiante UserType = 2
)
/* End Enum */

type Usuario struct {
	Username string
	Type     UserType
}

func main() {
	eduardo := Usuario {Username: "Eduardo", Type: Estudiante}
	uriel := Usuario {Username: "Uriel", Type: Maestro}

	fmt.Println(eduardo)
	fmt.Println(uriel)

	if eduardo.Type == Estudiante{
		fmt.Println("El usuario",eduardo.Username,"es un estudiante")
	}
}
