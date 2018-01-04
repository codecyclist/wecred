package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"m9d.de/wecred/features"
	"m9d.de/wecred/scoring"
)

func gatherFeatures(doc *goquery.Document) scoring.FeaturesSet {
	return scoring.FeaturesSet{
		LinkCount:           features.ComputeLinkCount(doc),
		WordFrequencyResult: features.ComputeWordFrequency(doc),
		SiteTitlesResult:    features.ComputeSiteTitle(doc),
	}
}

func executeScorings(featureSet scoring.FeaturesSet) (score scoring.Score) {
	score.BaseScore = 10
	for _, provider := range scoring.ScoreProviders {
		score.Scorings = append(score.Scorings, provider(featureSet)...)
	}
	return
}

func main() {
	config, confError := BuildAppConfig()

	if confError != nil {
		log.Fatal(confError)
		os.Exit(1)
	}

	doc, err := goquery.NewDocument(config.Url)

	if err != nil {
		log.Fatal(err)
	} else {
		featuresSet := gatherFeatures(doc)
		score := executeScorings(featuresSet)

		if config.Verbose {
			for _, v := range score.Scorings {
				fmt.Println(v.Pretty())
			}
		}

		fmt.Printf("Score=%d of %d\n", score.GetRating(), score.BaseScore)
	}

}
