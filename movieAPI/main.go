package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Structs for the API
type DB struct {
	collection *mongo.Collection
}

type Movie struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Year      string             `json:"year" bson:"year"`
	Directors []string           `json:"directors" bson:"directors"`
	Writers   []string           `json:"writers" bson:"writers"`
	BoxOffice BoxOffice          `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross  uint64 `json:"gross" bson:"gross"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("appDB").Collection("movies")
	db := &DB{collection: collection}

	// Create a router and define the routes
	router := mux.NewRouter()
	router.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}",
		db.GetMovie).Methods(http.MethodGet)
	router.HandleFunc("/v1/movies",
		db.PostMovie).Methods(http.MethodPost)
	router.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}",
		db.DeleteMovie).Methods(http.MethodDelete)
	router.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}",
		db.PutMovie).Methods(http.MethodPut)

	// Start the server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// GetMovie fetches a movie with a given ID
func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside GetMovie")
	vars := mux.Vars(r)
	var movie Movie
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	log.Println("ObjectID: ", objectID)
	filter := bson.M{"_id": objectID}
	err := db.collection.FindOne(context.TODO(),
		filter).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		log.Println("Response: ", string(response))
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		log.Println("Response sent")
	}
}

// PostMovie adds a new movie to our MongoDB collection
func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside PostMovie")
	var movie Movie
	postBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(postBody, &movie)
	result, err := db.collection.InsertOne(context.TODO(), movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// DeleteMovie deletes a movie with a given ID
func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside DeleteMovie")
	// Get the id from the request
	vars := mux.Vars(r)
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{
		"_id": objectID,
	}
	// Delete the movie
	_, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Movie deleted")
}

// PutMovie updates a movie with a given ID
func (db *DB) PutMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside PutMovie")
	// Get the id param from the request
	vars := mux.Vars(r)
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	// Decode the incoming movie json
	var movie Movie
	postBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(postBody, &movie)
	// Prepare the update model
	update := bson.M{
		"$set": movie,
	}
	// Update the movie
	filter := bson.M{"_id": objectID}
	_, err := db.collection.UpdateOne(context.TODO(),
		filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Movie updated")
}
