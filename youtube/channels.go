package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	for{
		msg := <-channel
		fmt.Println("Estoy imprimiendo lo que recibí del canal:", msg)
	}
}
