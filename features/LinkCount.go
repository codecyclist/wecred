package features

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var affiliateLinksExpression, _ = regexp.Compile("(\\.amazon\\.)")
var maskedAffiliateLinksExpression, _ = regexp.Compile("(amzn\\.to)")
var shortenedUrlExpression, _ = regexp.Compile("(bit\\.ly|tinyurl\\.com)")

type LinkCountResult struct {
	TotalLinks           int `json:"TotalLinks"`
	LocalLinks           int `json:"LocalLinks"`
	AffiliateLinks       int `json:"AffiliateLinks"`
	MaskedAffiliateLinks int `json:"MaskedAffiliateLinks"`
	ShortendedUrls       int `json:"ShortendedUrls"`
}

func ComputeLinkCount(doc *goquery.Document) LinkCountResult {
	result := LinkCountResult{}
	localdomain := doc.Url.Host

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		result.TotalLinks++
		link, exists := s.Attr("href")

		if exists {
			if strings.Contains(link, localdomain) {
				result.LocalLinks++
			}

			if affiliateLinksExpression.MatchString(link) {
				result.AffiliateLinks++
			}

			if maskedAffiliateLinksExpression.MatchString(link) {
				result.MaskedAffiliateLinks++
			}

			if shortenedUrlExpression.MatchString(link) {
				result.ShortendedUrls++
			}
		}

	})
	return result
}
