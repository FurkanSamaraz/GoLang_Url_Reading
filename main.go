package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	respons, _ := http.Get("https://www.fullhdfilmizlesene.com/film")
	doc, _ := goquery.NewDocumentFromReader(respons.Body)

	doc.Find(".dty").Each(func(i int, s *goquery.Selection) {
		tex := s.Find("a").Text()
		fmt.Println(i+1, tex)
	})

}
