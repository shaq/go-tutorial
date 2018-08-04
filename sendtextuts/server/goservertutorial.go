package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa, Go is awesome!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Expert web design by Shaquizzle :)")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about/", about)
	http.ListenAndServe(":8000", nil)
}
