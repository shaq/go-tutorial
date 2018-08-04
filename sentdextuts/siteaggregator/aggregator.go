package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex represents <sitemap> tag of a sitemap.
// type SitemapIndex struct {
// 	Locations []string `xml:"sitemap>loc"`
// }

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
	// var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap)

	// file := "washingtonpost.xml"
	// washPostXML, err := os.Open(file)
	//
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// fmt.Println("Successfully opened", file)
	// defer washPostXML.Close()
	//
	// bytes, _ := ioutil.ReadAll(washPostXML)

	// xml.Unmarshal(bytes, &s)
	// for idx, loc := range s.Locations {
	// 	s.Locations[idx] = strings.TrimSpace(loc)
	// }

	// for _, Loc := range s.Locations {
	resp, err := http.Get("https://www.theguardian.com/sitemaps/news.xml")
	fmt.Println(err)
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	xml.Unmarshal(bytes, &n)

	for idx := range n.Titles {
		newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Location[idx]}
	}
	// }

	for idx, data := range newsMap {
		fmt.Println(idx, data.Keywords, data.Location)
	}
}
