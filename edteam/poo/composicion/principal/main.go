package main

import (
	"example.com/customer"
	"example.com/invoice"
	"example.com/invoiceitem"
	"fmt"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Popayán",
		customer.New("Alejandro", "Cl 123 #23-4", "9876543210"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "Curso de Go", 12.34),
			invoiceitem.New(2, "Curso de POO con Go", 54.23),
			invoiceitem.New(3, "Curso de Testin con Go", 90),
		),
	)
	i.SetTotal()

	fmt.Printf("%+v\n", i)
}
