package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func multi_returns() {
	fn, sn := getInitials("tifa lockhart")
	fmt.Println(fn, sn)

	fn1, sn1 := getInitials("tifa")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("tifa lockhart")
	fmt.Println(fn2, sn2)

}
