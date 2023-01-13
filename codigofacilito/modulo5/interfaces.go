package main

import "fmt"

type Animal interface {
	comer()
	dormir()
	cazar()
}

type Perro struct {
	Nombre string
}

func (self Perro) comer() {
	fmt.Println("El perro come!")
}

func (self Perro) dormir() {
	fmt.Println("El perro duerme!")
}

func (self Perro) cazar()  {
	fmt.Println("El perro caza!")
}

func ejecutarAcciones(animal Animal)  {
	animal.comer()
	animal.dormir()
	animal.cazar()
}

func main() {
	loki := Perro{Nombre: "Loki"}

	fmt.Println(loki)
	ejecutarAcciones(loki)
}
