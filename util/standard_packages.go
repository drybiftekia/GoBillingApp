package main

import (
	"fmt"
	"sort"
	"strings"
)

func standard_packages() {
	greeting := "hello there friends!"

	fmt.Println("This should be true:", strings.Contains(greeting, "hello"))
	fmt.Println("This should be false:", strings.Contains(greeting, "hello!"))

	fmt.Println("I can use in place replacements:", strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println("The index for 'th'", strings.Index(greeting, "th"))
	fmt.Println("Splitting the string on the spaces returns the array:", strings.Split(greeting, " "))

	fmt.Println("Previous operations did not change the greetings string:", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25, 30}

	sort.Ints(ages)
	fmt.Println("The sort packages changes the original variable", ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "bowser"))
}
