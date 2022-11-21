package main

import "fmt"

// Las constantes van fuera de la función main
const CURSO string = "Curso Profesional de Go!"

func main() {
	//var nombre, apellido, pais = "Omar","Martínez","México"
	nombre, apellido, pais := "Omar", "Martínez", "México"
	edad, altura := 44, 1.68

	fmt.Println(nombre, apellido, pais, edad, altura)
	fmt.Println(CURSO, "\n ")

	var booleano = edad > 5
	var nombre1 = "Cody"
	var nombre2 = "Cody "

	fmt.Println(booleano)
	fmt.Println(nombre1 == nombre2)
}
