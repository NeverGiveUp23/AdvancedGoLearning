package main

import (
	"fmt"
)

func main() {
	cart := []string{"apple", "orange"}

	for i := range cart {
		fmt.Println(cart[i])
	}

	fruit := cart[:2]
	fmt.Println("fruit:", fruit)

	out := concat([]string{"A", "B"}, []string{"C"})
	fmt.Println("concat:", out)
}

func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)

	return s
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) == cap(s) {
		size := 2 * (len(s) + 1)
		ns := make([]int, size)
		copy(ns, s)
		s = ns[:len(s)]
	}

	s = s[:len(s)+1]
	s[i] = v
	return s
}
