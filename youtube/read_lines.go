package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	archivo, error := os.Open("./hola.txt")

	if error != nil {
		fmt.Println("Hubo un error al abrir el archivo")
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
}
