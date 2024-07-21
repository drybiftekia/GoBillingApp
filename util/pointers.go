package main

import "fmt"

func updateName2(x *string) {
	*x = "wedge"
}

func pointers() {
	name := "tifa"

	// updateName(name)

	fmt.Println("memory address of name is:", &name)

	m := &name
	fmt.Println(m)
	fmt.Println("value at memory address:", *m)
	fmt.Println(name)
	updateName2(m)
	fmt.Println(name)
}
