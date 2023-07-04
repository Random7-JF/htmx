package main

import (
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Coppola"},
				{Title: "Blade Runner", Director: "Scott"},
				{Title: "The Thing", Director: "Carpenter"},
			},
		}
		template.Execute(w, films)

	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
