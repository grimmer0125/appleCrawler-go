package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// func (s *Selection) Each(f func(int, *Selection)) *Selection {
// 	for i, n := range s.Nodes {
// 		f(i, newSingleSelection(n, s.document))
// 	}
// 	return s
// }
func StartCrawer(handler func(macs []Mac)) {
	doc, err := goquery.NewDocument("http://www.apple.com/tw/shop/browse/home/specialdeals/mac")
	if err != nil {
		log.Fatal(err)
	}

	var macs []Mac

	// Find the review items
	doc.Find(".refurb-list .box-content table").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		// band := s.Find("a").Text()
		title := s.Find(".specs h3").Text()
		title = strings.TrimSpace(title)

		url, _ := s.Find(".specs h3 a").Attr("href")

		price := s.Find(".purchase-info .price").Text()
		price = strings.TrimSpace(price)

		// fmt.Printf("product %d: %s, %s,%s \n", i, title, url, price)
		// total := title + ";" + url + ";" + price

		mac := Mac{title, url, price}
		macs = append(macs, mac)
	})

	handler(macs)
}
