package main

import (
	"fmt"
	"sync"
)

func crawlerResult() *[]Information {
	informationList := startCrawlers()
	return informationList
}

func startCrawlers() *[]Information {
	var wg sync.WaitGroup
	wg.Add(1)
	go homePageCrawler(&wg)
	wg.Add(1)
	go weeklyCrawler(&wg)
	wg.Add(1)
	go monthlyCrawler(&wg)
	wg.Add(1)
	go yearlyCrawler(&wg)

	go func() {
		// Wait for all goroutines to finish
		wg.Wait()
		fmt.Println("channel closed")
		close(crawledInformationChannel)
	}()

	signMap := make(map[string]Information)

	for informationMap := range crawledInformationChannel {
		for sign, timeframeMap := range informationMap {
			info, ok := signMap[sign]
			if !ok {
				// If not, create a new Information struct
				info = Information{sign: sign}
			}
			for timeframe, text := range timeframeMap {
				switch timeframe {
				case "today":
					info.today = text
				case "weekly":
					info.weekly = text
				case "monthly":
					info.monthly = text
				case "yearly":
					info.year = text
				}
			}
			info.percDetails = generatePercentageDetails()
			// Update the signMap
			signMap[sign] = info
		}
	}
	for _, info := range signMap {
		informationList = append(informationList, info)
	}

	return &informationList
}
