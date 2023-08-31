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

func GreetAll(gs ...Greeter)  {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Valor: %v, Tipo: %T\n", g, g)
	}
}

func main()  {
	p := Person{Name: "Alejandro"}
	var t Text = "Daisy"
	// p.Greet()
	// t.Greet()
	GreetAll(p, t)
}