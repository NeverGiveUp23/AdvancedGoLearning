package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// given github login, return name of public repos
// make http req, then passing json needed



func main() {
	response , err := http.Get("https://api.github.com/users/nevergiveup23")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// always check the status is okay
	if response.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s\n", response.Status)
	}

	// getting the content type 
	ctype := response.Header.Get("Content-Type")
	fmt.Println("Content-Type:", ctype)
	
	var reply struct {
		Name        string `json:"name"`
		NumRepos int    `json:"public_repos"`
	}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&reply); err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(reply.Name, reply.NumRepos)


	// io.Copy -> copies from the reader to the writer which is the Stdout
	// io.Copy just reads the incoming data from the request body (response.Body) in this case and sends somewhere else (os.Stdout) in this case

}

/* JSON <-> Go
   Types:
   string <-> string
	true/false <-> bool
	number <-> float, float32, int, int32, int8 ..., uint, uint8 ...
	array <-> []T. []any
	object <-> map[string]any, struct
	

	encodingjson api
	JSON <-> []byte -> go: Unmarshal
	Go -> []byte -> JSON: Marshal
	JSON -> io.Reader -> Go: Decoder
Go -> io.Writer -> JSON: Encoder 

*/
