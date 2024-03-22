package main

import "fmt"

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	//Tipo de dato
	fmt.Printf("Tipo de datos de baseCuadrado: %T\n", baseCuadrado)

	// Devolución de dos valores
	value1, value2 := doubleReturn(8)
	fmt.Println("Valor1: ", value1, " Valor2: ", value2)

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
	entry := []string{"Jack","John","Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 0 {
			break;
		}
	}

}
