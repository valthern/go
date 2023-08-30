package main

import "fmt"

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola, soy %s\n", p.Name)
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola, soy %s\n", t)
}

func main()  {
	// var g Greeter = Person{ Name: "Alejandro"}
	var g Greeter = Text("Daisy")
	g.Greet()
}