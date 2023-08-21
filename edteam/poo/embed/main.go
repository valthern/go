package main

type Person struct {
	Name string
	Age	 uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet()  {
	
}