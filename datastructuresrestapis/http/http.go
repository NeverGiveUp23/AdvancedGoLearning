package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// given github login, return name of public repos
// make http req, then passing json needed

type GitHubClient struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
}

func NewGitHubClient() *GitHubClient {
	return &GitHubClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL:   "https://api.github.com",
		userAgent: "your-app-name/1.0",
	}
}

func (g *GitHubClient) GetUserInfo(ctx context.Context, login string) (string, int, error) {
	if login == "" {
		return "", 0, fmt.Errorf("login cannot be empty")
	}

	url := g.baseURL + "/users/" + login
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", g.userAgent)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := g.httpClient.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("failed to fetch user data: %w", err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		return parseResponse(res.Body)
	case http.StatusNotFound:
		return "", 0, fmt.Errorf("user %q not found", login)
	case http.StatusForbidden:
		return "", 0, fmt.Errorf("Github api rate limit exceeded")
	default:
		return "", 0, fmt.Errorf("Github api error for %q: %s", login, res.Status)
	}
}

func main() {
	// always check the status is okay
	//	var githubUsername string
	//	fmt.Println(userInfo(githubUsername))
	// io.Copy -> copies from the reader to the writer which is the Stdout
	// io.Copy just reads the incoming data from the request body (response.Body) in this case and sends somewhere else (os.Stdout) in this case

	fmt.Println("Enter Github username: ")
	var login string
	if _, err := fmt.Scanln(&login); err != nil {
		fmt.Printf("Error reading input %v\n", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := NewGitHubClient()
	result, count, err := client.GetUserInfo(ctx, login)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s, Count: %d\n", result, count)
}

// userInfo() returns the name and number of public_repos from GitHub API
/*

func userInfo(login string) (string, int, error) {
	if _, err := fmt.Scan(&login); err != nil {
		return "", 0, err
	}
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


*/

func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Name     string `json:"name"`
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
