package scoring

import (
	"strings"
)

func SuspiciousTitleScore(featuresSet FeaturesSet) (scorings []Scoring) {
	if strings.Contains(featuresSet.SiteTitlesResult.URL.Host, "-test") {
		scorings = append(scorings, Scoring{
			Issuer: SuspiciousTitleScore,
			Points: -1,
			Title:  "Suspicious URL",
			Remark: "Pretends to deliver tests, which mostly means affiliate links",
		})
	}
	return
}

func init() {
	register(SuspiciousTitleScore)
}
