package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mehulsuthar-000/find-mirrors/mirrors"
)

type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		response := findFastest(mirrors.MirrorList)
		respJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJson)
	})

	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL)
			latency := time.Since(start)
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()
	}
	return response{<-urlChan, <-latencyChan}
}
