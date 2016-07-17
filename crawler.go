package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func StartCrawer() ([]Mac, error) {
	doc, err := goquery.NewDocument("http://www.apple.com/tw/shop/browse/home/specialdeals/mac")
	if err != nil {
		log.Fatal(err)
	}

	var macs []Mac

	doc.Find(".refurb-list .box-content table").Each(func(i int, s *goquery.Selection) {

		title := s.Find(".specs h3").Text()
		title = strings.TrimSpace(title)

		url, _ := s.Find(".specs h3 a").Attr("href")

		price := s.Find(".purchase-info .price").Text()
		price = strings.TrimSpace(price)

		// fmt.Printf("product %d: %s, %s,%s \n", i, title, url, price)

		mac := Mac{title, url, price}
		macs = append(macs, mac)
	})

	return macs, nil
}
