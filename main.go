package main

import (
	"MTS/auth/httpserver"
	"log"
)

func main() {
	server, err := httpserver.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.S.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
