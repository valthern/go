package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	archivo, error := os.Open("./hola.txt")
	defer func ()  {
		archivo.Close()
		fmt.Println("Defer")
	}()
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

	if true {
		return true
	}

	fmt.Println("Nunca me ejecuto")
	
	return true;
}
