package main

import "fmt"

type Exampler interface {
	x()
}

func wrapper(i interface{}) {
	fmt.Printf("valor: %v, Tipo: %T\n", i, i)
}

func main() {
	// var e Exampler
	// e.x()
	wrapper(12)
	wrapper(14.32)
	wrapper(false)
	wrapper("Alejandro")
}
