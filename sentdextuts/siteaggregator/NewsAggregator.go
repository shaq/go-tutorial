package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

// NewsAggPage presents the page for the news aggregator.
type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

// Sitemapindex represents the main site sitemap which lists
// all other sitemaps.
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
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

func newsRoutine(c chan News, Location string) {

}

func newsAgg(w http.ResponseWriter, r *http.Request) {
	var s Sitemapindex
	var n News

	newsMap := make(map[string]NewsMap)
	resp, err := http.Get("https://www.telegraph.co.uk/sitemap.xml")
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	xml.Unmarshal(bytes, &s)

	for _, Loc := range s.Locations {
		resp, _ := http.Get(Loc)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		resp.Body.Close()

		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Location[idx]}
		}
	}

	p := NewsAggPage{Title: "Shaq's News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("newsaggtemplate.html")
	fmt.Println(t.Execute(w, p))
}
