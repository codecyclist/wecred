package scoring

import (
	"fmt"

	"github.com/codecyclist/wecred/features"
)

type Score struct {
	BaseScore int
	Scorings  []Scoring
}

func (s Score) GetRating() (score int) {
	score = s.BaseScore
	for _, v := range s.Scorings {
		score += v.Points
	}
	if score < 0 {
		score = 0
	}
	return
}

type Scoring struct {
	Points int
	Title  string
	Remark string
	Issuer ScoreProvider
}

func (s Scoring) Pretty() string {
	return fmt.Sprintf("Penalty points=%d reason=%s details=%s", s.Points, s.Title, s.Remark)
}

type FeaturesSet struct {
	LinkCount           features.LinkCountResult
	WordFrequencyResult features.WordFrequencyResult
	SiteTitlesResult    features.SiteTitlesResult
}

type ScoreProvider func(featuresSet FeaturesSet) (scorings []Scoring)

var ScoreProviders = make([]ScoreProvider, 0)

func register(provider ScoreProvider) {
	ScoreProviders = append(ScoreProviders, provider)
}
