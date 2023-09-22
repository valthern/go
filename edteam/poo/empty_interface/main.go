package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}

func wrapper(i interface{}) {
	// v, ok := i.(string)
	// if ok {
	// 	fmt.Printf("\t%s\n", strings.ToUpper(v))
	// }

	switch v := i.(type) {
	case string:
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Printf("valor: %v, Tipo: %T\n", i, i)
	}
}

func main() {
	wrapper(12)
	wrapper("AdeAdalberto")
	wrapper(34.68595)
	wrapper(false)
	wrapper("Qué pedo")
}
