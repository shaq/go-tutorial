package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type News struct {
	Titles   []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Location []string `xml:"url>loc"`
}

type NewsMap struct {
	Keywords string
	Location string
}

func main() {
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

	for idx, data := range newsMap {
		fmt.Println(idx, data.Keywords, data.Location)
	}
}
