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

	err = uploadDetails(ctx, client)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(client)
}

func uploadDetails(ctx context.Context, client *firestore.Client) error {
	data := *prepareData()

	for key, value := range data {
		_, err := client.Collection("details").Doc(key).Set(ctx, value)
		if err != nil {
			log.Fatalf("error setting document: %v\n", err)
		}
	}

	return nil
}
