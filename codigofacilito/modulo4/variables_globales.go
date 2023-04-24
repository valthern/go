package main

import "fmt"

var username string 

func funcion1()  {
	username = "Funcion 1"
	fmt.Println("Función1:",username)
}

func funcion2()  {
	username = "Funcion 2"
	fmt.Println("Función2:",username)
}

func main()  {
	username = "Cody"
	funcion1()
	funcion2()

	fmt.Println(username)
}