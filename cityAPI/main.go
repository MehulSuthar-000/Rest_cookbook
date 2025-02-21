package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	Name string `json:"name"`
	Area uint64 `json:"area"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Create a new city instance
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		fmt.Printf("Got %s city with area of %d sq miles!\n",
			tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	http.HandleFunc("/city", postHandler)
	http.ListenAndServe(":8080", nil)
}
