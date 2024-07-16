package main

import "fmt"

func loops() {
	x := 0
	// for can be for and while loop
	for x < 5 { // while loop
		fmt.Println("The value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++ { // for loop
		fmt.Println("The value of i is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the position at index %v, %v\n", index, value)
	}

	for _, value := range names { // NOTE: This does not alter the value rather just the local value in the function
		fmt.Printf("the value is %v\n", value)
		value = "new string"
	}

	fmt.Println(names)
}
