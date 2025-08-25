package main

import (
	"fmt"
	"sync"
)

func channelMake(c chan int, x int) {
	c <- x
	close(c)
}

func printer(ch chan bool) {
	ch <- true
}

func main() {
	ch := make(chan int, 1)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func(ch chan int) {
		defer waitGroup.Done()
		channelMake(ch, 42)
		fmt.Println("Exit...")
	}(ch)
	fmt.Println("Read...", <-ch)

	_, ok := <-ch
	if ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("Channel is closed!")
	}

	waitGroup.Wait()
	var c chan bool = make(chan bool)
	for i := 0; i < 5; i++ {
		go printer(c)
	}

	n := 0
	for i := range c {
		fmt.Println(i)
		if i == true {
			n++
		}
		if n > 2 {
			fmt.Println("n:", n)
			close(c)
			break
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

}
