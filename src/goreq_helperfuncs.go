package main

import (
	"fmt"
	"io"
	"net/http"
)

func http_request_get() {

	var url_for_get_request string

	fmt.Print("What URL/API would you like to do a GET request to?: ")

	fmt.Scanln(&url_for_get_request)

	get_response, err := http.Get(url_for_get_request)

	if err != nil {

		fmt.Println("Error initiating GET request", err)
		return
	}

	defer get_response.Body.Close()

	formatted_body, err := io.ReadAll(get_response.Body)

	if err != nil {

		fmt.Println("Couldn't read the response body", err)
		return
	}

	fmt.Println(string(formatted_body))

}

func http_request_head() {

	var url_for_head_request string

	fmt.Print("What URL/API would you like to do a HEAD request to?: ")

	fmt.Scanln(&url_for_head_request)

	get_response, err := http.Head(url_for_head_request)

	if err != nil {

		fmt.Println("Error initiating HEAD request", err)
		return
	}

	defer get_response.Body.Close()

	formatted_body, err := io.ReadAll(get_response.Body)

	if err != nil {

		fmt.Println("Couldn't read the response body", err)
		return
	}

	fmt.Println(string(formatted_body))

}

func http_request_post() {

	var url_for_post_request string

	fmt.Print("What URL/API would you like to do a POST request to?: ")

	defer fmt.Println("WARNING: No data was actually sent!")

	fmt.Scanln(&url_for_post_request)

	get_response, err := http.Post(url_for_post_request, "text/plain", nil)

	if err != nil {

		fmt.Println("Error initiating POST request", err)
		return
	}

	defer get_response.Body.Close()

	formatted_body, err := io.ReadAll(get_response.Body)

	if err != nil {

		fmt.Println("Couldn't read the response body", err)
		return
	}

	fmt.Println(string(formatted_body))

}
