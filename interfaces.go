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
	if user.Permisos() > 4 {
		return user.Name() + " tiene permisos de administrador"
	}
	return ""
}

func main() {
	admin := Admin{"Uriel"}
	editor := Editor{"marcos"}

	fmt.Println(auth(admin))
	fmt.Println(auth(editor))
}
