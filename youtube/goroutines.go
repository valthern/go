package main

import(
	"fmt"
	"time"
	"strings"
)

func main()  {
	go mi_nombre_lentooo("Estereotipo")
	fmt.Println("Queeee aburridoooooo")
	
	go func ()  {	
		var wait string
		fmt.Scanln(&wait)
	}()
}

func mi_nombre_lentooo(name string)  {
	letras := strings.Split(name,"")

	for _,letra := range(letras){
		time.Sleep(time.Second)
		fmt.Println(letra)
	}
}
