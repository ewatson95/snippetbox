package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port = ":3000"

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

func handleSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display snippet with id: %d", id)
}

func handleSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", handleHome)
	mux.HandleFunc("GET /snippet/view", handleSnippetView)
	mux.HandleFunc("POST /snippet/create", handleSnippetCreate)

	log.Printf("Starting on port %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
