package main

import (
	"log"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleHome)

	log.Print("Starting on port :3000")
	err := http.ListenAndServe("localhost:3000", mux)
	log.Fatal(err)
}
