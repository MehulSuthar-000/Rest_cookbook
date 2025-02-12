package main

import (
	"io"
	"net/http"
	"time"
)

// Create Handlerfor HealthCheck API
// HealthCheck API returns date time to client
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}

func main() {
	// Register HealthCheck API Handler
	http.HandleFunc("/healthCheck", HealthCheck)

	// Start Server
	http.ListenAndServe(":8080", nil)
}
