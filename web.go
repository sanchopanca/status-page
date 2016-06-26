package main

import (
	"html/template"
	"net/http"
)

func routes() {
	http.HandleFunc("/", index)
}

type Person struct {
	UserName string
}

var templ = template.Must(template.ParseFiles("html/page.html"))

func index(w http.ResponseWriter, r *http.Request) {
	p := Person{UserName: "User"}
	templ.ExecuteTemplate(w, "page.html", p)
}

func main() {
	routes()
	http.ListenAndServe(":8080", nil)
}
