package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	executeReadFile()
	fmt.Println("Nunca me voy a imprimir")
}

func executeReadFile()  {
	ejecucion := read_File()
	fmt.Println(ejecucion)
}

func read_File() bool {
	archivo, error := os.Open("./holaa.txt")

	defer func() {
		archivo.Close()
		fmt.Println("Defer")

		recover()
	}()

	if error != nil {
		fmt.Println("Hubo un error al abrir el archivo")
		panic(error)
	}

	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Printf("%d %s\n", i, linea)
		//fmt.Println(linea)
	}
	fmt.Println()

	if true {
		return true
	}

	fmt.Println("Nunca me ejecuto")

	return true
}
