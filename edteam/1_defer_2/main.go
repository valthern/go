package main

import (
	"fmt"
	"os"
)

func main()  {
	file, err := os.Create("hello.txt")

	if err != nil {
	fmt.Printf("Ocurrió un error al crear el archivo %v\n",err)
		return
	}

	_, err = file.Write([]byte("Hello EDteam"))
	if err != nil {
		fmt.Printf("Ocurrió un erro al escribir el archivo: %v\n",err)
		return
	}
}