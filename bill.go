package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 4.99},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breadkown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		// Spacing added and the ':' is added through the value
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}
