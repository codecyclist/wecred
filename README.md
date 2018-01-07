# Website Crediblity Score Generator
This tool computes the expected crediblity of a website, primarily by assuming that websites who use affiliate marketing are in a conflict of interest or even not supposed to provide any real information at all.

## Why did I write this tool?
Basically, I wrote this tool for two reasons.

First, I wanted to avoid accidentally opening Affiliate Garbage pages.

Second, I needed a pretty showcase for a Golang talk.

## What does the Name mean?
Some say that wecred is short for "Website Credibility", others might state that it slightly sounds like "regret", which could refer to regretting having wasted time reading a shallow glibberish article on an affiliate website.


## Scoring
A website has an ideal score of 10. Wecred will apply a brief scan of the content and issue penalty points for the following disturbances:

* Affiliate links (1 Penalty Point)
* Many Affiliate links (2 Penalty Point)
* Disguised Affiliate Links, trying to fool readers (3 Penalty Points)
* Many Disguised Affiliate Links (4 Penalty Points)
* Keyword which is suspect to be SEO (1 Penalty Point)
* Suspicious Title tags (1 Penalty Point)

(subject to be extended with further updates)

So, the lower the score, the lower the expected credibility of the site. From experience, reading sites with a score lower than 6 might not be superior to other ways to invest time.

## Usage

Commandline interface:
```
wecred --verbose http://example.com
```

Example Output (with actual URL and keywords removed)

```
wecred --verbose http://(censored)-test.com/                                                     
Penalty points=-1 reason=Has Affiliate Links details=
Penalty points=-2 reason=Has Multiple Affiliate Links details=
Penalty points=-3 reason=Has Masked Affiliate Links details=
Penalty points=-4 reason=Has Multiple Masked AffiliateLinks details=
Penalty points=-1 reason=Possible Keyword details=word=xxxx occourences=11 average=3
Penalty points=-1 reason=Possible Keyword details=word=yyyyyyy occourences=46 average=3
Penalty points=-1 reason=Possible Keyword details=word=zzzzzzz occourences=17 average=3
Penalty points=-1 reason=Possible Keyword details=word=aaaaaaa occourences=21 average=3
Penalty points=-1 reason=Possible Keyword details=word=bbbbb occourences=10 average=3
Penalty points=-1 reason=Suspicious URL details=Pretends to deliver tests, which mostly means affiliate links
Score=0 of 10

```
