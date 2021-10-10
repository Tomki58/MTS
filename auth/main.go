package main

import (
	"auth/httpserver"
	"log"
)

func main() {
	server, err := httpserver.New("")
	if err != nil {
		log.Fatal(err)
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
