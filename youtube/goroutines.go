package main

import(
	"fmt"
	"time"
	"strings"
)

func main()  {
	go mi_nombre_lentooo("Omar")
	fmt.Println("Queeee aburridoooooo")
}

func mi_nombre_lentooo(name string)  {
	letras := strings.Split(name,"")

	for _,letra := range(letras){
		time.Sleep(1 * time.Second)
		fmt.Println(letra)
	}
}
