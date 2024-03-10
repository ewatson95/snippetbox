package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	files := []string{"./ui/html/base.tmpl", "./ui/html/pages/home.tmpl"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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
