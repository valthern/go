package main

import "fmt"

func main()  {
	// Declaramos el mapa sin usar la función Make()
	usuarios := map[int] string {}

	usuarios[1] = "Usuario 1"
	usuarios[2] = "Usuario 2"
	usuarios[3] = "Usuario 3"
	usuarios[4] = "Usuario 4"
	usuarios[5] = "Usuario 5"
	usuarios[6] = "Usuario 6"
	usuarios[7] = "Usuario 7"
	usuarios[8] = "Usuario 8"
	usuarios[9] = "Usuario 9"

	for llave, valor := range usuarios {
		fmt.Println(llave, valor)
	}
}