package main

import "fmt"

func main() {
	willClose := make(chan complex64, 10)

	willClose <- -1
	willClose <- 1i

	<-willClose
	<-willClose
	close(willClose)

	read := <-willClose
	fmt.Println(read)
}

func printer(ch chan<- bool) {
	ch <- true
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
