package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}

func wrapper(i interface{}) {
	fmt.Printf("valor: %v, Tipo: %T\n", i, i)

	v, ok := i.(string)
	if ok {
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	}
}

func main() {
	wrapper(12)
	wrapper("AdeAdalberto")
	wrapper(34.68595)
	wrapper(false)
	wrapper("Qué pedo")
}
