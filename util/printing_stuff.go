package main

import "fmt"

func printing_stuff() {

	name := "Bob"
	age := 15

	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	fmt.Println("hello guys!")
	fmt.Println("Goodbye guys!")
	fmt.Println("my age is age", age, "and my name is", name)

	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("you scored %0.2f congratulations.\n", 250.5)

	// Save formated strings
	var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println("the saved string is:", str)
}
