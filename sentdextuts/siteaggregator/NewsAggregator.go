package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// NewsAggPage presents the page for the news aggregator.
type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

// News represents the Titles, Keywords, and Locations of a
// given sitemap.
type News struct {
	Titles   []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Location []string `xml:"url>loc"`
}

// NewsMap is a struct representing each individual News
// article.
type NewsMap struct {
	Keywords string
	Location string
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa, Go is awesome!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Expert web design by Shaquizzle :)")
}

func newsAgg(w http.ResponseWriter, r *http.Request) {
	var n News
	newsMap := make(map[string]NewsMap)

	resp, err := http.Get("https://www.theguardian.com/sitemaps/news.xml")
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	xml.Unmarshal(bytes, &n)

	for idx := range n.Titles {
		newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Location[idx]}
	}

	p := NewsAggPage{Title: "Shaq's News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("newsaggtemplate.html")
	fmt.Println(t.Execute(w, p))
}
