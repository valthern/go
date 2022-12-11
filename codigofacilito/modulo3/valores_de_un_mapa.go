package main

import (
	"fmt"
)

func main() {
	usuarios := make(map[string]string)

	// usuarios["eduardo"] = "eduardo@codigofacilito.com"
	usuarios["eduardo"] = ""

	// En un mapa se obtienen dos valores: el valor y un booleano, que indica si existe el valor o no
	if correo, ok := usuarios["eduardo"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
