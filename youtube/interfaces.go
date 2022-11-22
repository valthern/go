package main

import "fmt"

type User interface {
	Permisos() int
	Name() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}
func (this Admin) Name() string {
	return this.nombre
}

type Editor struct {
	nombre string
}

func (this Editor) Permisos() int {
	return 3
}
func (this Editor) Name() string {
	return this.nombre
}

func auth(user User) string {
	permisos := user.Permisos()

	if permisos > 4 {
		return user.Name() + " tiene permisos de administrador"
	} else if permisos == 3 {
		return user.Name() + " tiene permisos de editor"
	}
	return ""
}

func main() {
	// admin := Admin{"Uriel"}
	// editor := Editor{"Marcos"}

	usuarios := []User{Admin{"Uriel"},Editor{"Fulano"}}

	for _,usuario := range usuarios{
		fmt.Println(auth(usuario))
	}

	// fmt.Println(auth(admin))
	// fmt.Println(auth(editor))
}
