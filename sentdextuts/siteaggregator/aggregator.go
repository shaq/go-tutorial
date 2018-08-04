package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var washPostXML = []byte(`
  <sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-politics-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-politics-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-opinions-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-opinions-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-local-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-local-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-sports-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-sports-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-national-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-national-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-world-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-world-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-business-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-business-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-technology-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-technology-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-lifestyle-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-lifestyle-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-entertainment-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-entertainment-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-blogs-goingoutguide-sitemap.xml
  </loc>
  </sitemap>
  <sitemap>
  <loc>
  http://www.washingtonpost.com/news-goingoutguide-sitemap.xml
  </loc>
  </sitemap>
  </sitemapindex>
`)

// SitemapIndex represents <sitemap> tag of a sitemap.
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles   []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Locs     []string `xml:"url>loc"`
}

func main() {
	bytes := washPostXML

	var s SitemapIndex
	var n News
	xml.Unmarshal(bytes, &s)

	for _, loc := range s.Locations {
		resp, _ := http.Get(loc)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		fmt.Println(loc)
	}
}
