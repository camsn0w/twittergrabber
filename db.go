package twittergrabber

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var ctx mongo.SessionContext

func getClient() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://admin:smalltoast20@cluster0.ubr9l.mongodb.net/<dbname>?retryWrites=true&w=majority",
	))
	if err != nil {
		println("Cannot create client")
		log.Fatal(err)
	}
	return client

}
