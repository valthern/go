package main

import "fmt"

type UseR struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User UseR
	Id   string
}

func main() {
	eduardo := UseR{Name: "Eduardo", Email: "eduardo@codigo.com", Active: true}
	uriel := UseR{Name: "Uriel", Email: "uriel@codigo.com", Active: true}

	eduardoStudent := Student{User: eduardo,Id: "1f1f1f1f1"}

	fmt.Println(uriel)
	fmt.Println(eduardoStudent.User.Active)
}
