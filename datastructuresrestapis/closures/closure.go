package main

import "fmt"

func main() {
	seq := adder()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
}

// Function that takes no arguments but returns another function that returns an int
func adder() func() int {
	i := 0
	fmt.Println("value of i: ", i)

	return func() int {
		i++
		fmt.Println("added 1 to i: ", i)
		return i
	}
}
