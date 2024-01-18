package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.URLFilters(
			regexp.MustCompile(`(?m)google.com`),
		),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		logger.Info("onHTML", zap.Any("e", e.Attr("href")))
		// e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("The status code ", r.StatusCode)
		logger.Info("response", zap.Any("headers", r.Headers))
	})

	c.Visit("https://google.com/")
}
