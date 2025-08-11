package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// given github login, return name of public repos
// make http req, then passing json needed

func main() {
	// always check the status is okay
	fmt.Println(userInfo("nevergiveup23"))
	// io.Copy -> copies from the reader to the writer which is the Stdout
	// io.Copy just reads the incoming data from the request body (response.Body) in this case and sends somewhere else (os.Stdout) in this case

}
// userInfo() returns the name and number of public_repos from github api
func userInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login

	res, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	if res.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%q - bad status: %s", url, res.Status)
	}
	return parseResponse(res.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	
	var reply struct {
		Name        string `json:"name"`
		NumRepos int    `json:"public_repos"`
	}

	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&reply); err != nil {
		fmt.Println("ERROR: ", err)
		return "", 0, err
	}
	return reply.Name, reply.NumRepos, nil
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
