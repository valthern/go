package main

import (
	"fmt"
	"time"
)

func hi(num int) {
	fmt.Println("Hola ", num)
	time.Sleep(1 * time.Second)
}

func main() {
	for i := 0; i < 10; i++ {
		go hi(i)
	}

	var s string
	fmt.Scan(&s)
}
