package main

import (
	"fmt"

	"example.com/course"
)

func main() {
	Go := course.New("Go desde Cero", 12.34, false)

	Go.SetName("POO con Go")
	Go.SetPrice(89.25)
	Go.SetIsFree(false)
	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	fmt.Println(Go.Name())
	fmt.Println(Go.Price())
	fmt.Println(Go.IsFree())
	Go.UserIDs()
	Go.Classes()

	fmt.Printf("\n%+v\n", Go)
}
