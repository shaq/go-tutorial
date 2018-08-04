package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex represents <sitemap> tag of a sitemap.
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location represents <loc> tag of sitemap.
type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.theverge.com/sitemaps")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)
	for _, loc := range s.Locations {
		fmt.Println(loc)
	}
}
