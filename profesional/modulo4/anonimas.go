package main

import "fmt"

func main() {
	miFuncion := func(username string) string{
		message := fmt.Sprintf("Hola %s, te saludamos desde una función sin nombre",username)

		return message
	}

	mensajeFinal := miFuncion("Cody 1")

	fmt.Println(mensajeFinal)
}
