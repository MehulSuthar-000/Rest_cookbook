package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/mehulsuthar-000/urlShortener/data"
)

// Input data
// create a struct that will represent the input data
type ShortenRequest struct {
	Url string `json:"url"`
}

func main() {
	// define router
	// we will use the gorilla/mux router
	router := mux.NewRouter()

	// define routes
	// we will use the HandleFunc method to define routes
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/shorten", shortenHandler)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/v1/{url}", redirectHandler)

	// configure server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// start server
	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}

// create a shortener function that will return a short link
// how it works:
// 1. generate a short link
// 2. save the short link in the database
// 3. return the short link
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shorten handler called")
	// parse input data
	// we will use the ShortenRequest struct to parse the input data
	var ShortenRequest ShortenRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ShortenRequest)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		fmt.Println("Error decoding request: ", err)
		return
	}
	// generate a short link
	// we will use the generateShortLink function to generate a short link
	shortLink := generateShortLink(ShortenRequest.Url)

	// save the short link in the database
	data.Db[shortLink] = ShortenRequest.Url

	// return the short link
	// we will return the short link in the JSON format
	response := map[string]string{
		"shortLink": shortLink,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}

// generate a short link
func generateShortLink(url string) string {
	// generate a short link based on the URL
	// for simplicity, we will use the first 5 characters of the URL
	return url[:5]
}

// create a redirect function that will redirect to the original link
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirect handler called")
	vars := mux.Vars(r)

	// Get the short link from the URL
	shortLink := vars["url"]

	// log the database
	fmt.Printf("Database: %#v\n", data.Db)

	// Get the original link from the database
	originalLink, ok := data.Db[shortLink]
	if !ok {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	if !strings.HasPrefix(originalLink, "http://") && !strings.HasPrefix(originalLink, "https://") {
		originalLink = "http://" + originalLink // Default to HTTP if missing
	}
	http.Redirect(w, r, originalLink, http.StatusSeeOther)

}
