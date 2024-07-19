package main

import "fmt"

func maps() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		262435923: "mario",
		123931589: "luigi",
		129192492: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[262435923])

	phonebook[262435923] = "bowser"
	fmt.Println(phonebook)

}
