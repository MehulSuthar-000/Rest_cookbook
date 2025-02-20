// Contains the handlers for the sender entity
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/MehulSuthar-000/LogisticAPI/database"
	"github.com/MehulSuthar-000/LogisticAPI/models"
	"github.com/MehulSuthar-000/LogisticAPI/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var SenderCollection *mongo.Collection

func InitCollections() {
	// ❌ Check if `database.Database` is nil before using it
	if database.Database == nil {
		log.Panicln("❌ Database is nil! MongoDB not connected before InitCollections.")
	}

	SenderCollection = database.Database.Collection("sender")

	// ✅ Debug log
	if SenderCollection == nil {
		log.Panicln("❌ SenderCollection is still nil. Something went wrong!")
	} else {
		fmt.Println("✅ SenderCollection initialized successfully!")
	}
}

func GetSender(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Sender"))
}

func GetAllSenders(w http.ResponseWriter, r *http.Request) {

	// fetch sender collection from db
	var SenderCollection = database.GetCollection("sender")

	// fetch sender from db
	result, err := SenderCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Panicln("Error fetching senders: ", err)
		http.Error(w, "Error fetching senders", http.StatusInternalServerError)
		return
	}

	// Since we are fetching all the senders, we will store them in a slice
	senders := []models.Sender{}

	// iterate over the cursor and decode the data
	for result.Next(context.Background()) {
		var sender models.Sender
		result.Decode(&sender)
		senders = append(senders, sender)
	}

	// send the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.ToJSON(senders))
}

func CreateSender(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Sender"))
}

func UpdateSender(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Sender"))
}

func DeleteSender(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Sender"))
}
