package main

import "fmt"

type course struct {
	name string
}

func (c course) Print()  {
	fmt.Printf("%+v\n",c)
}

// declaración de alias
type myAlias = course

// definición de tipo
type newCourse course

func main()  {
	c := newCourse{name: "Go"}
	// c.Print()
	fmt.Printf("El tipo es: %T\n", c)
}
