package main

import "fmt"

type User struct{
	Name string
	Email string
}

func (self *User) setName(name string)  {
	self.Name = name	
}

func main()  {
	cody := User{Name: "Cody", Email: "info@codigo.com"}

	fmt.Println()
	
}