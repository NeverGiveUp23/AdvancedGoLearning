package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}

func gen(min, max int, createNumber chan int, end chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case createNumber <- rand.IntN(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
			// return
		case <-time.After(4 * time.Second):
			fmt.Println("time.After()!")
			return
		}
	}
}

// Fibinacci can take a long time to run in the program so lets make a function spinner using a go routine to wait for the functions end
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
