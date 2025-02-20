package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MehulSuthar-000/LogisticAPI/database"
	"github.com/MehulSuthar-000/LogisticAPI/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// initialize mongo db connection
	database.ConnectDB()

	// initialize collections
	handlers.InitCollections()

	// Create router
	router := mux.NewRouter()

	// Create routes for each API

	// sender API
	router.HandleFunc("/sender", handlers.CreateSender).Methods("POST")
	router.HandleFunc("/sender", handlers.GetAllSenders).Methods("GET")
	router.HandleFunc("/sender/{id}", handlers.GetSender).Methods("GET")
	router.HandleFunc("/sender/{id}", handlers.UpdateSender).Methods("PUT")
	router.HandleFunc("/sender/{id}", handlers.DeleteSender).Methods("DELETE")

	// receiver API
	router.HandleFunc("/receiver", handlers.CreateReceiver).Methods("POST")
	router.HandleFunc("/receiver", handlers.GetAllReceivers).Methods("GET")
	router.HandleFunc("/receiver/{id}", handlers.GetReceiver).Methods("GET")
	router.HandleFunc("/receiver/{id}", handlers.UpdateReceiver).Methods("PUT")
	router.HandleFunc("/receiver/{id}", handlers.DeleteReceiver).Methods("DELETE")

	// shipment API
	// router.HandleFunc("/shipment", handlers.CreateShipment).Methods("POST")
	// router.HandleFunc("/shipment", handlers.GetShipments).Methods("GET")
	// router.HandleFunc("/shipment/{id}", handlers.GetShipment).Methods("GET")
	// router.HandleFunc("/shipment/{id}", handlers.UpdateShipment).Methods("PUT")
	// router.HandleFunc("/shipment/{id}", handlers.DeleteShipment).Methods("DELETE")

	// // Carrier API
	// router.HandleFunc("/carrier", handlers.CreateCarrier).Methods("POST")
	// router.HandleFunc("/carrier", handlers.GetCarriers).Methods("GET")
	// router.HandleFunc("/carrier/{id}", handlers.GetCarrier).Methods("GET")
	// router.HandleFunc("/carrier/{id}", handlers.UpdateCarrier).Methods("PUT")
	// router.HandleFunc("/carrier/{id}", handlers.DeleteCarrier).Methods("DELETE")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
