package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

/*func greet(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello World %s", time.Now())
*/

func main() {

	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)

	//buffered reader
	reader := bufio.NewReader(strings.NewReader("Hello, from bufio package!!\n"))

	//reading byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data) // Transfer data from one point to another -> source to the target
	if err != nil {
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, data[:n]) // Data[:n] -> limit up to the last char since we can only hold 20 bytes

	// readString -> allows you to continue reading from the NewReader as a continuation
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string", err)
		return
	}

	fmt.Println("Read string: ", line)

	writer := bufio.NewWriter(os.Stdout)
	// Writing byte slice
	data1 := []byte("Hello bufio package\n")
	num, err := writer.Write(data1)
	if err != nil {
		return
	}
	// we need to fluch the data in the buffer
	defer func() {
		if err := writer.Flush(); err != nil {
			log.Printf("Flush Error: %v", err)
		}
	}()

	fmt.Printf("Wrote %d bytes\n", num)

}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye %s", time.Now())
}
