package main

import (
	"strconv"
	"time"
)

type HomePage struct {
	TodayDescription string `json:"today_description"`
	TodayHealth      string `json:"today_health"`
	TodayLove        string `json:"today_love"`
	TodayMoto        string `json:"today_moto"`
	TodayCareer      string `json:"today_career"`
	YearDescription  string `json:"year_description"`
	YearHealth       string `json:"year_health"`
	YearLove         string `json:"year_love"`
	YearMoto         string `json:"year_moto"`
	YearCareer       string `json:"year_career"`
}

func newHomePage(information Information) *HomePage {
	return &HomePage{
		TodayDescription: information.today,
		TodayHealth:      information.percDetails.TodayHealth,
		TodayLove:        information.percDetails.TodayLove,
		TodayMoto:        information.percDetails.TodayMoto,
		TodayCareer:      information.percDetails.TodayCareer,
		YearDescription:  information.year,
		YearHealth:       information.percDetails.YearHealth,
		YearLove:         information.percDetails.YearLove,
		YearMoto:         information.percDetails.YearlyMoto,
		YearCareer:       information.percDetails.YearCareer,
	}
}

func mapInformationToHomePage(informationList *[]Information) map[string]map[string]HomePage {

	currentDay := strconv.Itoa(time.Now().Day())
	signToHomePage := make(map[string]HomePage)

	for _, info := range *informationList {
		homePage := *newHomePage(info)
		signToHomePage[info.sign] = homePage
	}
	signToIndex := make(map[string]string)
	for sign, index := range zodiacSigns {
		signToIndex[sign] = index
	}

	signToHomePagesByIndex := make(map[string]map[string]HomePage)
	for sign, homepage := range signToHomePage {
		index, ok := signToIndex[sign]
		if !ok {
			continue // Skip signs not found in the zodiacSigns map
		}
		temp := make(map[string]HomePage)

		temp[currentDay] = homepage

		signToHomePagesByIndex[index] = temp
	}
	return signToHomePagesByIndex
}

func prepareFirebaseHomePage(informationList *[]Information) *map[string]map[string]HomePage {
	data := mapInformationToHomePage(informationList)
	return &data
}
