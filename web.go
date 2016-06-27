package main

import (
	"html/template"
	"net/http"
)

func routes() {
	http.HandleFunc("/", index)
}


var templ = template.Must(template.ParseFiles("html/page.html"))

func index(w http.ResponseWriter, r *http.Request) {
	var statuses = make(map[string]string)
	statuses["DB"] = "OK"
	statuses["WEB"] = "FAIL"
	statuses["API"] = "OK"
	templ.ExecuteTemplate(w, "page.html", statuses)
}

func main() {
	routes()
	http.ListenAndServe(":8080", nil)
}
