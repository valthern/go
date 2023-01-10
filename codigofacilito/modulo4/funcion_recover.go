package main

import "fmt"
import "os"

func main() {

	defer func ()  {
		if err := recover(); err != nil{
			fmt.Println("Ups, al parecer el programa no finalizó de forma correcta!")
		}
	}()

	if file, err := os.Open("ejemplo.txt"); err != nil {
		panic("No fue posible obtener el archivo!")
	} else {
		defer func() {
			fmt.Println("Cerramos el archivo!")
			file.Close()
		}()

		contenido := make([]byte, 254)
		longitud, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[0:longitud])

		fmt.Println(contenidoArchivo)
	}
}
