package main

import (
	"errors"
	"fmt"
	//"io/ioutil"
	"strings"
)

func main() {
	/* 	texto := "oMaR"
	   	minus, mayus := convert(texto)
	   	fmt.Println(minus, mayus)

	   	content, err := ioutil.ReadFile("./things.txt")
	   	if err != nil {
	   		fmt.Printf("Ocurrio un error: %v\n", err)
	   		return
	   	}

	   	fmt.Println(string(content))
	*/

	result, err := division(10, 2)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}

	fmt.Println(result)
}

func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}

func division(dividendo, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("No puedes dividor entre cero\n")
		return
	}

	result = dividendo / divisor
	return
}
