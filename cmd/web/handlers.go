package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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
