package main

import (
	"log"
	"net/http"
)

// func readinessCheck(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Content-Type: text/plain; charset=utf-8"))
// }

func main() {

	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot))))

	newServer := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {

		w.Header().Set("Content-Type:", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	if err := newServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
