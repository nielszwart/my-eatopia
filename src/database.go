package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var user User

func initDb() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

func loginUser(email string, password string) bool {
	client, ctx := initDb()
	database := client.Database("myeatopia")
	userCollection := database.Collection("user")
	defer client.Disconnect(ctx)

	filter := bson.D{{"email", email}, {"password", password}}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return false
	}

	templateData.User = user

	return true
}

func getRecipesForUser(user int) []*Recipe {
	client, ctx := initDb()
	database := client.Database("myeatopia")
	collection := database.Collection("recipe")

	var recipes []*Recipe
	filter := bson.D{{"UserID", user}}
	resultsCursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	for resultsCursor.Next(ctx) {
		var recipe Recipe
		err := resultsCursor.Decode(&recipe)
		if err != nil {
			log.Fatal(err)
		}

		recipes = append(recipes, &recipe)
	}

	if err := resultsCursor.Err(); err != nil {
		log.Fatal(err)
	}

	resultsCursor.Close(ctx)

	return recipes
}
