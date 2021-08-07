package main

import (
	"github.com/gocolly/colly"
)

func scrap(word string) []string {
	c := colly.NewCollector()

	var res []string

	c.OnHTML("span.one-click-content", func(e *colly.HTMLElement) {
		res = append(res, e.Text)
	})

	url := "https://www.dictionary.com/browse/" + word
	c.Visit(url)

	return res
}
