package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

const remote = "http://localhost:3000/login"

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		Transport: nil,
		Jar:       jar,
		Timeout:   0,
	}

	request, err := http.NewRequest("GET", "http://localhost:3000/login", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.SetBasicAuth("Aladdin", "open sesame")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
