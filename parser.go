package main

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"errors"
)

type Parser interface {
	parse(address string) []string
}

type WebsiteParser struct {}

func(wp WebsiteParser) parse(address string) ([]string, error) {
	var links []string
	resp, err := http.Get(address)
	defer resp.Body.Close()
	if err == nil {
		if resp.StatusCode == 404 {
			return nil, errors.New("wiki page does not exist")
		} else if resp.StatusCode != 200 {
			return nil, errors.New("an unexpected error occurred")
		}
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		doc.Find("a").Each(func(i int, selection *goquery.Selection) {
			l, _ := selection.Attr("href")
			links = append(links, l)
		})
	} else {
		return nil, err
	}
	return links, nil
}
