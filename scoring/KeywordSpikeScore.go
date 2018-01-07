package scoring

import "fmt"

func KeywordSpikeScore(featuresSet FeaturesSet) (scorings []Scoring) {
	if wordCardinality := len(featuresSet.WordFrequencyResult.WordCounts); wordCardinality > 0 {
		wordAverageFrequency := featuresSet.WordFrequencyResult.TotalWords / wordCardinality

		for word, occourences := range featuresSet.WordFrequencyResult.WordCounts {
			if occourences > 3*wordAverageFrequency && len(word) > 6 {
				scorings = append(scorings, Scoring{
					Issuer: KeywordSpikeScore,
					Points: -1,
					Title:  "Possible Keyword",
					Remark: fmt.Sprintf("word=%s occourences=%d average=%d", word, occourences, wordAverageFrequency),
				})
			}
		}
	}
	return
}

func init() {
	register(KeywordSpikeScore)
}
