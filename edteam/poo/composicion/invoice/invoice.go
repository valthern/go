package invoice

import (
	"example.com/customer"
	"example.com/invoiceitem"
)

// Invoice is the structure of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// New returns a new invoice
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
