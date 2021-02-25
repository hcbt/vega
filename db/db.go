package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectToDB establishes Database connection
func ConnectToDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost/")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func AddEntry(entry string) {
	client := ConnectToDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("vega").Collection("tracks")

	track := Track {
		ID: 	entry,
		Artist: "",
		Title:  "",
		Genre:  "",
		Tags:   []string{},
	}

	_, err = collection.InsertOne(ctx, track)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
}

func FindEntry(value string) (Track, error) {
	var result Track

	filter := bson.D{{"_id", value}}

	client := ConnectToDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("vega").Collection("tracks")

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	defer client.Disconnect(ctx)

	fmt.Println("result:", result)

	return result, nil
}
