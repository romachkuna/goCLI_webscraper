package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	crawledInformationChannel = make(chan map[string]map[string]string)
	informationList           []Information
	zodiacSigns               = map[string]string{
		"Aries":       "0",
		"Taurus":      "1",
		"Gemini":      "2",
		"Cancer":      "3",
		"Leo":         "4",
		"Virgo":       "5",
		"Libra":       "6",
		"Scorpio":     "7",
		"Sagittarius": "8",
		"Capricorn":   "9",
		"Aquarius":    "10",
		"Pisces":      "11",
	}
)

type Information struct {
	sign        string
	today       string
	weekly      string
	monthly     string
	year        string
	percDetails percentageDetails
}

type NestedDetails struct {
	DescriptionDaily   string `json:"description_daily"`
	DescriptionWeekly  string `json:"description_weekly"`
	DescriptionMonthly string `json:"description_monthly"`
	DescriptionYearly  string `json:"description_yearly"`
	TodayCareer        string `json:"today_career"`
	TodayHealth        string `json:"today_health"`
	TodayLove          string `json:"today_love"`
	TodayMoto          string `json:"today_moto"`
	YearCareer         string `json:"year_career"`
	YearHealth         string `json:"year_health"`
	YearLove           string `json:"year_love"`
	YearlyMoto         string `json:"yearly_moto"`
}

type percentageDetails struct {
	TodayCareer string `json:"today_career"`
	TodayHealth string `json:"today_health"`
	TodayLove   string `json:"today_love"`
	TodayMoto   string `json:"today_moto"`
	YearCareer  string `json:"year_career"`
	YearHealth  string `json:"year_health"`
	YearLove    string `json:"year_love"`
	YearlyMoto  string `json:"yearly_moto"`
}

func generatePercentageDetails() percentageDetails {
	return percentageDetails{
		TodayCareer: strconv.Itoa(rand.Intn(51) + 50),
		TodayHealth: strconv.Itoa(rand.Intn(51) + 50),
		TodayLove:   strconv.Itoa(rand.Intn(51) + 50),
		TodayMoto:   "Today Moto",
		YearCareer:  strconv.Itoa(rand.Intn(51) + 50),
		YearHealth:  strconv.Itoa(rand.Intn(51) + 50),
		YearLove:    strconv.Itoa(rand.Intn(51) + 50),
		YearlyMoto:  "Yearly Moto",
	}
}

func mapInformationToDetails(informationList *[]Information) map[string]map[string]NestedDetails {

	currentDay := strconv.Itoa(time.Now().Day())

	signToDetails := make(map[string]NestedDetails)

	for _, info := range *informationList {
		nestedDetails := NestedDetails{
			DescriptionDaily:   info.today,
			DescriptionWeekly:  info.weekly,
			DescriptionMonthly: info.monthly,
			DescriptionYearly:  info.year,
			TodayCareer:        info.percDetails.TodayCareer,
			TodayHealth:        info.percDetails.TodayHealth,
			TodayLove:          info.percDetails.TodayLove,
			TodayMoto:          info.percDetails.TodayMoto,
			YearCareer:         info.percDetails.YearCareer,
			YearHealth:         info.percDetails.YearHealth,
			YearLove:           info.percDetails.YearLove,
			YearlyMoto:         info.percDetails.YearlyMoto,
		}
		signToDetails[info.sign] = nestedDetails
	}
	// Convert sign names to indexes
	signToIndex := make(map[string]string)
	for sign, index := range zodiacSigns {
		signToIndex[sign] = index
	}

	signToDetailsByIndex := make(map[string]map[string]NestedDetails)
	for sign, details := range signToDetails {
		index, ok := signToIndex[sign]
		if !ok {
			continue // Skip signs not found in the zodiacSigns map
		}
		temp := make(map[string]NestedDetails)

		temp[currentDay] = details

		signToDetailsByIndex[index] = temp
	}

	return signToDetailsByIndex
}

func prepareData() *map[string]map[string]NestedDetails {
	informationList := startCrawlers()
	fmt.Println("Crawlers have completed")
	data := mapInformationToDetails(informationList)
	return &data
}
