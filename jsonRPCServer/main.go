package main

import (
	jsonparse "encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Args struct to hold arguments passed to JSON-RPC methods
type Args struct {
	ID string
}

// Book struct holds Book JSON structure from the reply
type Book struct {
	ID     string `json:"id",omitempty`
	Name   string `json:"name",omitempty`
	Author string `json:"author",omitempty`
}

// JSONServer struct to hold the JSON-RPC methods
type JSONServer struct{}

func (j *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	// Read JSON File and load data
	absPath, _ := filepath.Abs("D:\\learn_go\\Rest_cookbook\\jsonRPCServer\\books.json")
	raw, readerr := os.ReadFile(absPath)
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}
	// Decode JSON raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}

	// Iterate over each book to find the given book
	for _, book := range books {
		if book.ID == args.ID {
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	// Create a new RPC Server
	server := rpc.NewServer()
	// Register the type of data requested as JSON
	server.RegisterCodec(json.NewCodec(), "application/json")
	// Register the JSONServer struct as the implementation of the "JSONServer" service
	server.RegisterService(new(JSONServer), "")

	router := mux.NewRouter()
	// Handle all requests to the path "/rpc" with the RPC server
	router.Handle("/rpc", server)
	http.ListenAndServe(":1234", router)
}
