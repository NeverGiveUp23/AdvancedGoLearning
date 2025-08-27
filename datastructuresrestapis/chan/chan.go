package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Channels are the connections between goroutines, It is a communication mechanism that lets on goroutine send values
// To another goroutine.
// Channels are a reference to the data structure created by make. When we copy a channel or pass one as an argument, we are copying a reference.
func main() {
	ch := make(chan int) // this is how you make a channel
	var x int
	ch <- x  // A send statement
	x = <-ch // A receive statement in an assignment statement
	<-ch     // A receive statement; result is discarded

	// Channels support a third operation, CLOSE() which sets a flag indicating that no more values will ever be sent on this channel.
	// Subsequent attempts to send will panic

	// Receive operations on a closed channelyeild the values that have been sent until no more values are left
	// Anymore receive operations thereafter complete immediately and yeild the zero value of the channels element type

	// You close the channel by calling the built in function:
	close(ch)

	// Unbuffered channel -> with no second aregument.
	ch = make(chan int)    // Unbuffered channel
	ch = make(chan int, 0) //  Unbuffered channel
	ch = make(chan int, 1) // Buffered channel with 1 capacity

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	// A send operation on an unbuffered channel blocks the sending goroutine until anopther goroutine executes a corresponding receive on the same channel, at which point the value is transmitted and both goroutines may continue
	// Conversely, if the receive operation was attempted first, the receiving goroutine is blocked until another goroutine performs a send on the same channel

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	mustopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
