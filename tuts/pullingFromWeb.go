package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SiteMapIndex struct {
	Locations []location `xml: "sitemap"`
}

type Location struct {
	Loc string `xml: "loc"`
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}
