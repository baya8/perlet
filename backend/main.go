package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:rapideact@localhost:27017"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	InsertOne(client)
}

func InsertOne(client *mongo.Client) {

	coll := client.Database("mongo").Collection("mongo")

	doc := bson.D{{"title", "sample record to insert"}}

	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		panic(err)
	}

	id := result.InsertedID
	fmt.Println(id)

}
