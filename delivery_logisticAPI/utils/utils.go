package utils

import (
	"encoding/json"
	"log"
)

func ToJSON(data interface{}) []byte {
	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling data to JSON")
		return nil
	}

	return jsonData
}
