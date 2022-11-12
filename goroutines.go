package main

import(
	"fmt"
	"time"
	"strings"
)

func main()  {
	mi_nombre_lentooo("Omar")
}

func mi_nombre_lentooo(name string)  {
	letras := strings.Split(name,"")

	for _,letra := range(letras){
		time.Sleep(1 * time.Second)
		fmt.Println(letra)
	}
}
