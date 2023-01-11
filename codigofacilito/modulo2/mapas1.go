package main

import (
	"fmt"
)

func main() {
	/* 	dias := make(map[int]string)

	   	dias[0] = "Domingo"
	   	dias[1] = "Lunes"
	   	dias[2] = "Martes"
	   	dias[3] = "Miércoles"
	   	dias[4] = "Jueves"
	   	dias[5] = "Viernes"
	   	dias[6] = "Sábado"

	   	dias[4] = "JUEVES"

	   	delete(dias, 4)

	   	fmt.Println(dias)
	*/
	usuarios := make(map[string][]int)

	usuarios["Eduardo"] = []int{10, 9, 8, 10, 5}

	fmt.Println(usuarios)
}
