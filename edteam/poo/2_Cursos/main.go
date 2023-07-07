package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func main() {
	Go := Course{
		"Go desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Css := Course{
		Name:   "CSS desde Cero",
		IsFree: true,
	}

	Js := Course{}
	Js.Name = "Curso JS"
	Js.UserIDs = []uint{12, 67}

	fmt.Println(Go.Name)
	fmt.Printf("%+v\n", Css)
	fmt.Printf("%+v\n", Js)
}
