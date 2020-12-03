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

	// データを登録
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "David",
		"last":  "Gilmour",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding %v", err)
	}

	// データの更新
	//_, err = client.Collection("users").Doc("LRrPFGWyXyg4sqhsKQEt").Set(ctx, map[string]interface{}{
	//	"first": "Nick",
	//	"last":  "Mason",
	//	"born":  1940,
	//})

	// データの読み取り
	//iter := client.Collection("users").Documents(ctx)
	//for {
	//	doc, err := iter.Next()
	//	if err == iterator.Done {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalf("Faild to iterate: %v", err)
	//	}
	//	fmt.Println(doc.Data())
	//}

	// 削除
	_, err = client.Collection("users").Doc("LRrPFGWyXyg4sqhsKQEt").Delete(ctx)

	if err != nil {
		log.Fatalf("Faild to iterate: %v", err)
	}

}
