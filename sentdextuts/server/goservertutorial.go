package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa, Go is awesome!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Expert web design by Shaquizzle :)")
}

func newsAgg(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Shaq's News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/agg/", newsAgg)
	http.ListenAndServe(":8000", nil)
}
