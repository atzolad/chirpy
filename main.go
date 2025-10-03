package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	newServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := newServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
