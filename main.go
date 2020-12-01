package main

import (
	"context"
	"google.golang.org/api/option"
	"log"

	firebase "firebase.google.com/go"
)

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace %v", err)
	}
}
