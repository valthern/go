package main

import (
	"fmt"

	"example.com/course"
)

func main() {
	Go := course.NewCourse("Go desde Cero", 0, false)
	Go.UserIDs: []uint{12, 56, 89}
	Go.Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		}
	}

	Go.PrintClasses()
	fmt.Println(Go.Price)
}
