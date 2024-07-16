package main

import "fmt"

func arrays_slices() {
	var ages [3]int = [3]int{20, 25, 30}

	names := [4]string{"yoshi", ",mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, "+Length:", len(ages))
	fmt.Println(names)

	// Slices
	var scores = []int{100, 50, 60}
	scores[2] = 200
	// Slices can be appended
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]
	rangeOne = append(rangeOne, "kooper")
	fmt.Println(rangeOne)

	rangeTwo := names[2:]
	fmt.Println(rangeTwo)

	rangeThree := names[:3]
	fmt.Println(rangeThree)
}
