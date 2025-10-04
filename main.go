package main

import (
	"log"
	"net/http"
)

func main() {

	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	newServer := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	if err := newServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
