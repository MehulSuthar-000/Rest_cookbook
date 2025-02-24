package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mehulsuthar-000/urlShortener/helper"
	base62 "github.com/mehulsuthar-000/urlShortener/utils"
)

const Addr = "127.0.0.1:8080"

type DBClient struct {
	db *sql.DB
}

func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)

	// Get the ID from base62 string
	id := base62.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow("SELECT url FROM web_url WHERE id = $1", id).Scan(&url)

	// Handle response details
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record

	postBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(postBody, &record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", record.URL).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]string{"encoded_string": base62.ToBase62(id)}
		log.Println("Generated ID from DB:", id)
		log.Println("Encoded ID:", base62.ToBase62(id))
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func main() {
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{
		db: db,
	}
	defer db.Close()

	// Create a new router
	router := mux.NewRouter()

	// Attach an elegant path with handler
	router.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]}",
		dbclient.GetOriginalURL).Methods(http.MethodGet)

	router.HandleFunc("/v1/short",
		dbclient.GenerateShortURL).Methods(http.MethodPost)

	server := &http.Server{
		Addr:    Addr,
		Handler: router,
		// Good Practice: Enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting the server...")
	log.Fatal(server.ListenAndServe())

}
