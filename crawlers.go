package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"sync"
	"time"
)

func signCrawler(absoluteURl string, wg *sync.WaitGroup, sign string, timeframe string) {
	defer wg.Done()
	crawler := colly.NewCollector(colly.AllowURLRevisit())
	crawler.SetRequestTimeout(120 * time.Second)

	crawler.OnHTML("div.horoscope-content", func(element *colly.HTMLElement) {
		element.ForEach("p", func(_ int, p *colly.HTMLElement) {
			if len(p.Text) > 30 {
				innerMap := make(map[string]string)
				innerMap[timeframe] = p.Text

				outerMap := make(map[string]map[string]string)
				outerMap[sign] = innerMap

				crawledInformationChannel <- outerMap

			}
		})
	})
	err := crawler.Visit(absoluteURl)
	if err != nil {
		fmt.Println("problem visiting - signCrawler")
	}
}

func homePageCrawler(wg *sync.WaitGroup) {
	defer wg.Done()
	crawlHomePage := colly.NewCollector(colly.AllowURLRevisit())
	crawlHomePage.SetRequestTimeout(120 * time.Second)

	crawlHomePage.OnHTML("div.grid.grid-6.grid-4-m", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, a *colly.HTMLElement) {
			sign := a.DOM.Find("h2").Text()
			link := a.Attr("href")
			absoluteURL := e.Request.AbsoluteURL(link)
			wg.Add(1)
			go signCrawler(absoluteURL, wg, sign, "today")
		})
	})

	err := crawlHomePage.Visit("https://www.sunsigns.com/")
	if err != nil {
		fmt.Println(err)
	}

}

func weeklyCrawler(wg *sync.WaitGroup) {
	defer wg.Done()
	crawlWeekly := colly.NewCollector(colly.AllowURLRevisit())
	crawlWeekly.SetRequestTimeout(120 * time.Second)
	crawlWeekly.OnHTML("div.sign-container.flex-center", func(weekly *colly.HTMLElement) {
		weekly.ForEach("a", func(_ int, a *colly.HTMLElement) {
			sign := a.DOM.Find("h3").Text()
			if sign != "" {
				link := a.Attr("href")
				absoluteURL := weekly.Request.AbsoluteURL(link)
				wg.Add(1)
				go signCrawler(absoluteURL, wg, sign, "weekly")
			}
		})
	})
	err := crawlWeekly.Visit("https://www.sunsigns.com/horoscopes/weekly")
	if err != nil {
		fmt.Println(err)
	}
}

func monthlyCrawler(wg *sync.WaitGroup) {
	defer wg.Done()
	monthlyCrawler := colly.NewCollector(colly.AllowURLRevisit())
	monthlyCrawler.SetRequestTimeout(120 * time.Second)
	monthlyCrawler.OnHTML("div.sign-container.flex-center", func(weekly *colly.HTMLElement) {
		weekly.ForEach("a", func(_ int, a *colly.HTMLElement) {
			sign := a.DOM.Find("h3").Text()
			if sign != "" {
				link := a.Attr("href")
				absoluteURL := weekly.Request.AbsoluteURL(link)
				wg.Add(1)
				go signCrawler(absoluteURL, wg, sign, "monthly")
			}
		})
	})
	err := monthlyCrawler.Visit("https://www.sunsigns.com/horoscopes/monthly")
	if err != nil {
		fmt.Println(err)
	}
}

func yearlyCrawler(wg *sync.WaitGroup) {
	defer wg.Done()
	yearlyCrawler := colly.NewCollector(colly.AllowURLRevisit())
	yearlyCrawler.SetRequestTimeout(120 * time.Second)
	yearlyCrawler.OnHTML("div.sign-container.flex-center", func(weekly *colly.HTMLElement) {
		weekly.ForEach("a", func(_ int, a *colly.HTMLElement) {
			sign := a.DOM.Find("h3").Text()
			if sign != "" {
				link := a.Attr("href")
				absoluteURL := weekly.Request.AbsoluteURL(link)
				wg.Add(1)
				go signCrawler(absoluteURL, wg, sign, "yearly")
			}
		})
	})
	err := yearlyCrawler.Visit("https://www.sunsigns.com/horoscopes/yearly")
	if err != nil {
		fmt.Println(err)
	}
}
