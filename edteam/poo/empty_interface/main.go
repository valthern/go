package main

import "fmt"

type exampler interface {
	x()
}

func wrapper(i interface{}) {
	fmt.Printf("valor: %v, Tipo: %T\n", i, i)
}

func main() {
	wrapper(12)
	wrapper(34.68595)
	wrapper(false)
	wrapper("Qué pedo")
}
