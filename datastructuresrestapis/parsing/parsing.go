package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://example.com:8080/path?query=param#fragment"
	parseURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing parseURL", err)
		return
	}

	fmt.Println("Sceme: ", parseURL.Scheme)
	fmt.Println("Host: ", parseURL.Host)
	fmt.Println("Port: ", parseURL.Port())
	fmt.Println("Path: ", parseURL.Path)
	fmt.Println("Raw Query: ", parseURL.RawQuery)
	fmt.Println("Fragment: ", parseURL.Fragment)

	rawURL1 := "https://example.com/path?name=John&age=30"

	parseURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing: ", err)
		return
	}

	queryParams := parseURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age: ", queryParams.Get("age"))

	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseURL.Query()
	query.Set("name", "John")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL: ", baseURL.String())

	values := url.Values{}

	values.Add("name", "John")
	values.Add("age", "30")

	encodedQuery := values.Encode()

	fmt.Println(encodedQuery)

}
