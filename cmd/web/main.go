package main

import (
	"log"
	"net/http"
)

const port = ":3000"

func main() {

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("GET /{$}", handleHome)
	mux.HandleFunc("GET /snippet/view", handleSnippetView)
	mux.HandleFunc("POST /snippet/create", handleSnippetCreate)

	log.Printf("Starting on port %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
