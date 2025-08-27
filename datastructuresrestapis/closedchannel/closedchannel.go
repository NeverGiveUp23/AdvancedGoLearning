package main

import (
	"fmt"
	"sync"
)

func main() {
	willClose := make(chan complex64, 10)

	willClose <- -1
	willClose <- 1i

	<-willClose
	<-willClose
	close(willClose)

	read := <-willClose
	fmt.Println(read)

	var wg sync.WaitGroup

	wg.Add(1)
	var ch chan bool = make(chan bool)
	// write 5 values to channel with a single goroutine
	go func(ch chan bool) {
		defer wg.Done()
		printer(ch, 5)
	}(ch)
	for val := range ch {
		fmt.Println(val, " ")
	}
	fmt.Println()
	for i := 0; i < 15; i++ {
		fmt.Println(<-ch, " ")
	}
	fmt.Println()
	wg.Wait()
}

func printer(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func writeToChannel(ch chan<- int, x int) {
	fmt.Println("1", x)
	ch <- x
	fmt.Println("2", x)
}

// accepts two channel parameters. However out is available for reading and in is available for writing.
func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (f2)", x)
	in <- x
	return
}
