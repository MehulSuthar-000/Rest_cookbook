package main

import (
	"fmt"
	"net/http"
)

// create a new middleware
func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")

		// pass control back to the original handler
		originalHandler.ServeHTTP(w, r)

		fmt.Println("Executing middleware after response phase!")
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler!")
	w.Write([]byte("OK"))
}

func main() {
	// HandlerFunc return a HTTP Handler
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalHandler))
	http.ListenAndServe(":8080", nil)
}
