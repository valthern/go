package main

import (
	"fmt"
	"time"
)

func hello() int {
	fmt.Println("Hola, comunidad EDteam 🤚")
	return 1
}

func main()  {
	// go hello()
	go func ()  {
		fmt.Println("Hola, comunidad EDteam desde función anónima 🖐")
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("Hola, Gophers 😎")
}