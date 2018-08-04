package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/agg/", newsAgg)
	http.ListenAndServe(":8000", nil)
}
