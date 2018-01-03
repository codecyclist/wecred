package features

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type SiteTitlesResult struct {
	TitleTag string  `json:"TitleTag"`
	URL      url.URL `json:"URL"`
}

func ComputeSiteTitle(doc *goquery.Document, results chan<- SiteTitlesResult) {
	result := SiteTitlesResult{}

	result.TitleTag = doc.Find("title").Contents().Text()
	result.URL = *doc.Url

	results <- result
}
