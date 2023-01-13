package main

import "fmt"

type User struct{
	Name string
	Email string
}

func (self *User) setName(name string)  {
	self.Name = name	
}

func (self *User) getName() string {
	return self.Name
}

func main()  {
	cody := User{Name: "Cody", Email: "info@codigo.com"}

	cody.setName("Augusto")

	fmt.Println(cody.getName())
	
}