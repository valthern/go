package main

import "fmt"
import "errors"

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir sobre cero.")
	} else {
		return dividendo / divisor, nil
	}
}

func main() {
	if resultado, err := division(10, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El resultado es:", resultado)
	}
}
