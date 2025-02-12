package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ArticleHandler is a simple HTTP handler function which writes the article
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Category: " + vars["category"] + "\n"))
	w.Write([]byte("ID: " + vars["id"] + "\n"))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
