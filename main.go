package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"m9d.de/wecred/features"
)

func main() {
	//done := make(chan bool)
	config, confError := BuildAppConfig()

	if confError != nil {
		log.Fatal(confError)
		os.Exit(1)
	}

	doc, err := goquery.NewDocument(config.Url)

	if err != nil {
		log.Fatal(err)
	}

	linkCount := make(chan features.LinkCountResult)
	siteTitles := make(chan features.SiteTitlesResult)
	wordFrequency := make(chan features.WordFrequencyResult)

	go features.ComputeLinkCount(doc, linkCount)
	//go features.ComputeSiteTitle(doc, siteTitles)
	//go features.ComputeWordFrequency(doc, wordFrequency)

	for {
		select {
		case linkCountResult := <-linkCount:
			fmt.Println("Links:", linkCountResult)

		case siteTitlesResult := <-siteTitles:
			fmt.Println("Titles:", siteTitlesResult)

		case wct := <-wordFrequency:
			fmt.Println("TOtal Words", wct.TotalWords, " count:", wct.CountedWords)
			for k, v := range wct.WordCounts {
				if v > 4 {
					fmt.Printf("%s %d \n\r", k, v)
				}
			}

		}
	}

}
