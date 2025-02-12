package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

// UUID is a custom multiplexer
type UUID struct {
}

func (u *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
}

func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%x", b)
}

func main() {
	mux := &UUID{}
	http.ListenAndServe(":8080", mux)
}
