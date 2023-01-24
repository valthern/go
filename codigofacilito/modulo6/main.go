package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func crearUsuario() {
	fmt.Println("Usuario creado exitosamente!")
}

func listarUsuarios() {
	fmt.Println("Listado de usuarios!")
}

func actualizarUsuario() {
	fmt.Println("Usuario actualizado exitosamente!")
}

func eliminarUsuario() {
	fmt.Println("Usuario eliminado exitosamente!")
}

func main() {
	var reader *bufio.Reader
	//var option string

	reader = bufio.NewReader(os.Stdin)

	fmt.Println("A.- Crear")
	fmt.Println("B.- Listar")
	fmt.Println("C.- Actualizar")
	fmt.Println("D.- Eliminar")

	fmt.Print("Ingresa una opción válida: ")
	option, err := reader.ReadString('\n')

	if err != nil {
		panic("No es posible obtener el valor!")
	}

	//option = strings.TrimSuffix(option, "\n")
	fmt.Print(option)
	
	switch option {
	case "a", "crear":
		crearUsuario()
	case "b", "listar":
		listarUsuarios()
	case "c", "actualizar":
		actualizarUsuario()
	case "d", "eliminar":
		eliminarUsuario()
	default:
		fmt.Println("Opción no válida!")
	}

	//fmt.Print(option)
	fmt.Println("----------")
}
