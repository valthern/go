package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "Bla bla bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (this Tutor) hablar() string {
	return this.Human.hablar() + " Bienvenidos a CódigoFacilito"
}

func main() {
	tutor := Tutor{Human{"Uriel"}, Dummy{"Ariel"}}

	//fmt.Println(tutor.name)
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.hablar())
}
