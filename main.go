package main

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("C:/Users/Romac/Downloads/zodiacapp-workingdir-firebase-adminsdk-ngvgb-d84513cd0b.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	result := crawlerResult()

	err = uploadDetails(ctx, client, result)
	if err != nil {
		log.Fatalln(err)
	}

	err = uploadHomePage(ctx, client, result)
	if err != nil {
		log.Fatalln(err)
	}

	//err = uploadCompatiblity(ctx, client)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(client)
}

func uploadDetails(ctx context.Context, client *firestore.Client, informationList *[]Information) error {
	data := *prepareFirebaseDetails(informationList)

	for key, value := range data {
		_, err := client.Collection("details").Doc(key).Set(ctx, value, firestore.MergeAll)
		if err != nil {
			log.Fatalf("error setting document: %v\n", err)
		}
	}

	return nil
}

func uploadHomePage(ctx context.Context, client *firestore.Client, informationList *[]Information) error {
	data := *prepareFirebaseHomePage(informationList)

	for key, value := range data {
		_, err := client.Collection("home_page_content").Doc(key).Set(ctx, value, firestore.MergeAll)
		if err != nil {
			log.Fatalf("error setting document: %v\n", err)
		}
	}
	return nil
}

//func uploadCompatiblity(ctx context.Context, client *firestore.Client) error {
//	dummyData := CompatibilityDetails{
//		CareerWorkDisc:   "",
//		CareerWorkPerc:   "",
//		FriendshipDisc:   "",
//		FriendshipPerc:   "",
//		MarriageDisc:     "",
//		MarriagePerc:     "",
//		RelationshipDisc: "",
//		RelationshipPerc: "",
//	}
//	dummyMap := make(map[string]CompatibilityDetails)
//	uploadMap := make(map[string]map[string]CompatibilityDetails)
//
//	for i := 0; i < 12; i++ {
//		dummyMap[strconv.Itoa(i)] = dummyData
//	}
//
//	for i := 0; i < 12; i++ {
//		uploadMap[strconv.Itoa(i)] = dummyMap
//	}
//
//	for key, value := range uploadMap {
//		_, err := client.Collection("compatibility").Doc(key).Set(ctx, value, firestore.MergeAll)
//		if err != nil {
//			log.Fatalf("error setting document: %v\n", err)
//		}
//	}
//	return nil
//
//}
