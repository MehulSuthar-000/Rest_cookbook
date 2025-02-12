package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", queryHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	srv.ListenAndServe()
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Got the following query params: \n"))
	for key, value := range queryParams {
		w.Write([]byte(key + ": " + value[0] + "\n"))
	}
}
