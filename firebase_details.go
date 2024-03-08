package main

import (
	"strconv"
	"time"
)

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

func prepareFirebaseDetails(informationList *[]Information) *map[string]map[string]NestedDetails {
	data := mapInformationToDetails(informationList)
	return &data
}
