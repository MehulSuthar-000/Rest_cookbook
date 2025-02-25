package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"example.com/mehulsuthar-000/jsonStore/helper"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type DBClient struct {
	db *gorm.DB
}

type PackageResponse struct {
	Package helper.Package `json:"Package"`
}

func main() {
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}

	defer db.Close()

	// Initialize the router and assign the handler func to it
	router := mux.NewRouter()

	router.HandleFunc("/v1/package/{id:[a-zA-Z0-9]}",
		dbclient.GetPackage).Methods(http.MethodGet)
	router.HandleFunc("/v1/package",
		dbclient.PostPackage).Methods(http.MethodPost)
	router.HandleFunc("/v1/package",
		dbclient.GetPackageByWeight).Methods(http.MethodGet)

	// Create a Server Struct and Listen to the Requests
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// Implement methods for each paths/http method

func (driver *DBClient) GetPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	vars := mux.Vars(r)
	// Handle response details
	driver.db.First(&Package, vars["id"])
	var PackageData interface{}
	// Unmarshal JSON string to interface
	json.Unmarshal([]byte(Package.Data), &PackageData)
	var response = PackageResponse{Package: Package}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJson, _ := json.Marshal(response)
	w.Write(respJson)
}

func (driver *DBClient) GetPackageByWeight(w http.ResponseWriter, r *http.Request) {
	var packages []helper.Package
	weight := r.FormValue("weight")
	// Handle the response details
	var query = "select * from \"Package\" where data->>'weight'=?"
	driver.db.Raw(query, weight).Scan(&packages)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJson, _ := json.Marshal(packages)
	w.Write(respJson)
}

func (driver *DBClient) PostPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	postBody, _ := io.ReadAll(r.Body)
	Package.Data = string(postBody)
	driver.db.Save(&Package)
	responseMap := map[string]interface{}{"id": Package.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}
