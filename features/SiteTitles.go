package features

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type SiteTitlesResult struct {
	TitleTag string  `json:"TitleTag"`
	URL      url.URL `json:"URL"`
}

func ComputeSiteTitle(doc *goquery.Document) SiteTitlesResult {
	return SiteTitlesResult{
		TitleTag: doc.Find("title").Contents().Text(),
		URL:      *doc.Url,
	}
}
