package main

import "fmt"

type User struct{
	Name string
	Email string
}

func main()  {
	// var cody User

	// cody.Name = "Cody"
	// cody.Email = "email@codigo.com"

	// cody := User {Name: "Cody",Email: "email@codigo.com"}

	Email := "email@codigo.com"
	Name := "Cody"

	cody := User {Name, Email}

	fmt.Println(cody.Name)
	fmt.Println(cody.Email)
}