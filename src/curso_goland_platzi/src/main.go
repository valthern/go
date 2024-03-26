package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	// Devolución de dos valores
	value1, value2 := doubleReturn(8)
	fmt.Println("Valor1: ", value1, " Valor2: ", value2)

	// Tipo de dato
	fmt.Printf("Tipo de datos de valor1: %T\n", value1)

	// Ciclos
	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For while
	counter := 0
	for counter < 10 {
		fmt.Printf("%d ", counter)
		counter++
	}
	fmt.Println()

	// For de rango
	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}

	// For forever
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// 	if counterForever > 0 {
	// 		break
	// 	}
	// }

	// if con 2 variables internas
	if valueI, errI := strconv.Atoi("69"); errI != nil {
		log.Fatal(errI)
	} else {
		fmt.Println("Valor I:", valueI)
	}

	// if con 2 varaiables externas
	valueE, errE := strconv.Atoi("99")
	if errE != nil {
		log.Fatal(errE)
	}
	fmt.Println("Valor E:", valueE)

	// switch para iterar sobre una variable
	switch hora := time.Now().Hour(); hora {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		fmt.Println("Buenos días. Son las", hora, "horas")
	case 12, 13, 14, 15, 16, 17, 18:
		fmt.Println("Buenas tardes. Son las", hora, "horas")
	case 19, 20, 21, 22, 23:
		fmt.Println("Buenas noches. Son las", hora, "horas")
	}

	// switch para iterar sobre evaluaciones con una variable
	gradosCentigradosDelAgua := 80
	switch {
	case gradosCentigradosDelAgua >= 100:
		fmt.Println("El agua está en estado gaseoso:", gradosCentigradosDelAgua, "°C")
	case gradosCentigradosDelAgua <= 0:
		fmt.Println("El agua está en estado sólido:", gradosCentigradosDelAgua, "°C")
	default:
		fmt.Println("El agua está en estado líquido:", gradosCentigradosDelAgua, "°C")
	}

	// Array
	var array [4]int
	fmt.Println("Array:", array)

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("Slice:", slice, "Length Slice:", len(slice), "Capacity Slice:", cap(slice))

	// Métodos en el slice (primer índice es inclusivo, el segudo es exclusivo)
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
