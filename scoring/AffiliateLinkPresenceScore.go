package scoring

func ProvideAffiliateLinkPresenceScore(featuresSet FeaturesSet) (scorings []Scoring) {
	if featuresSet.LinkCount.AffiliateLinks > 0 {
		scorings = append(scorings, Scoring{
			Issuer: ProvideAffiliateLinkPresenceScore,
			Points: -1,
			Title:  "Has Affiliate Links",
		})
	}

	if featuresSet.LinkCount.AffiliateLinks > 1 {
		scorings = append(scorings, Scoring{
			Issuer: ProvideAffiliateLinkPresenceScore,
			Points: -2,
			Title:  "Has Multiple Affiliate Links",
		})
	}

	if featuresSet.LinkCount.MaskedAffiliateLinks > 0 {
		scorings = append(scorings, Scoring{
			Issuer: ProvideAffiliateLinkPresenceScore,
			Points: -3,
			Title:  "Has Masked Affiliate Links",
		})
	}

	if featuresSet.LinkCount.MaskedAffiliateLinks > 1 {
		scorings = append(scorings, Scoring{
			Issuer: ProvideAffiliateLinkPresenceScore,
			Points: -4,
			Title:  "Has Multiple Masked AffiliateLinks",
		})
	}

	return
}

func init() {
	register(ProvideAffiliateLinkPresenceScore)
}
