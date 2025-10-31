package main

import "fmt"

func main() {

	fmt.Println("Welcome to goreq!")
	fmt.Println("Use it like this please: https://example.com")
	fmt.Print("What type of HTTP/HTTPS request would you like to make? GET/HEAD/POST: ")

	var http_request_type string
	fmt.Scan(&http_request_type)

	switch http_request_type {

	case "GET":
		http_request_get()
	case "HEAD":
		http_request_head()
	case "POST":
		http_request_post()
	default:
		fmt.Println("Unknown HTTP request type")
	}
}
