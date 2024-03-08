package main

import (
	"log"
	"net/http"
)

const port = ":3000"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", handleHome)
	mux.HandleFunc("GET /snippet/view", handleSnippetView)
	mux.HandleFunc("POST /snippet/create", handleSnippetCreate)

	log.Printf("Starting on port %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
