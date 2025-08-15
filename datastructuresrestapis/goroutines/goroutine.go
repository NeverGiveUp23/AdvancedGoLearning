package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// GOMAXPROCS -> allows you to check and configure cores available on you machine
	fmt.Print("You are using", runtime.Compiler)
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0)) // displays core amount on your machine -> 0 does not change the setting

	// creating a multiple  goroutine
	// using waitgroup to rid of time.Sleep() -> waitgroup allows us to Add(), Done(), Wait() our goroutine
	var waitGroup sync.WaitGroup
	// fmt.Printf("%v\n", waitGroup) // learning the waitGroup structure -> DO NOT DO

	count := 15
	fmt.Printf("Going to run multiple %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	// fmt.Printf("%v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting....")
}

func printme(x int) {
	fmt.Println(x)
}
