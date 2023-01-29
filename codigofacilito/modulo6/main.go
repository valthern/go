package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int
var users map[int]User

func crearUsuario() {
	clearConsole()

	fmt.Print("Ingresa un valor para username: ")
	username := readLine()
	fmt.Print("Ingresa un valor para email: ")
	email := readLine()
	fmt.Print("Ingresa un valor para edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir de un string a un entero.")
	}

	id++
	user := User{id, username, email, age}
	users[id] = user

	fmt.Println("Usuario creado exitosamente!")
}

func listarUsuarios() {
	clearConsole()

	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}
	fmt.Println("---")
}

func actualizarUsuario() {
	fmt.Println("Usuario actualizado exitosamente!")
}

func eliminarUsuario() {
	fmt.Println("Usuario eliminado exitosamente!")
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor!")
	} else {
		return strings.TrimSuffix(option, "\n")
	}
}

func main() {
	var option string
	users = make(map[int]User)
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A.- Crear")
		fmt.Println("B.- Listar")
		fmt.Println("C.- Actualizar")
		fmt.Println("D.- Eliminar")

		fmt.Print("Ingresa una opción válida: ")
		option = readLine()

		if option == "quit" || option == "q" {
			break
		}

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
	}

	fmt.Println("¡Adiós! ----------")
}
