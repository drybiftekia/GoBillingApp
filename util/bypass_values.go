package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func bypass_values() {
	name := "tifa"
	name = updateName(name)
	fmt.Println(name)

	menu := map[string]float64{
		"pe":        1.99,
		"ice cream": 6.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
