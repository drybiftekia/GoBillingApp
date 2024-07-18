package main

import "fmt"

var Points = []int{20, 90, 100, 45, 70}

func SayHello(n string) {
	fmt.Println("hello", n)
}

func ShowScore() {
	fmt.Println("You scored this many points:", Points)
}
