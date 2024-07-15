package main

import "fmt"

func basic_vars() {

	fmt.Println("--- Code was Executed ---")

	// strings are always double quotes
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println("Strings 1:", nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println("Strings 2:", nameOne, nameTwo, nameThree)

	// This should be only done the first time when initializing the string variable
	// NOTE: This can't be used outside of a function!
	nameFour := "yoshi"

	fmt.Println("Strings 3:", nameOne, nameTwo, nameThree, nameFour)

	// ints
	var ageOne int = 30
	var ageTwo = 30
	ageThree := 40

	fmt.Println("Integers 1:", ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint = 25

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 84483124812394321942392349.2
	// inferred always as float64
	scoreThree := 1.5

	fmt.Println("Bits & Memory:", numOne, numTwo, numThree, scoreOne, scoreTwo, scoreThree)

}
