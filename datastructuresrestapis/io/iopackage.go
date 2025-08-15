package main

import (
	"fmt"
	"io"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return
	}

	fmt.Println("Bytes: ", string(buf[:n]))

}

func main() {

}
