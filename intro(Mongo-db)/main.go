package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Movie holds a movie data
type Movie struct {
	Name      string   `bson:"name"json:"name"`
	Year      string   `bson:"year" json:"year"`
	Directors []string `bson:"directors" json:"directors"`
	Writers   []string `bson:"writers" json:"writers"`
	BoxOffice `bson:"boxOffice" json:"boxOffice"`
}

// BoxOffice is nested in Movie
type BoxOffice struct {
	Budget uint64 `bson:"budget" json:"budget"`
	Gross  uint64 `bson:"gross" json:"gross"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB successfully")
	collection := client.Database("appDB").Collection("movies")

	log.Println("Creating a new movie in MongoDB")

	// Create a movie
	darkNight := Movie{
		Name:      "The Dark Knight",
		Year:      "2008",
		Directors: []string{"Christopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffice: BoxOffice{
			Budget: 185000000,
			Gross:  533316061,
		},
	}
	// Insert a document into MongoDB
	_, err = collection.InsertOne(context.TODO(), darkNight)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Inserted a new movie in MongoDB")
	}

	queryResult := &Movie{}
	// bson.M is used for building map for filter query
	filter := bson.M{"boxOffice.budget": bson.M{"$gt": 150000000}}
	result := collection.FindOne(context.TODO(), filter)
	err = result.Decode(queryResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie:", queryResult)

	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Disconnected from MongoDB")
}
