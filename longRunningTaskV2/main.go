package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

const (
	queueName  string = "jobQueue"
	hostString string = "127.0.0.1:8000"
)

func HandleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

/*
Creates a server object and initiates
the Channel and Queue details to publish messages
*/
func getServer(name string) JobServer {
	conn, err := amqp.Dial("amqp://guest:guest@172.20.36.69:5672/")
	HandleError(err, "Dialing failed to RabbitMQ broker")
	channel, err := conn.Channel()
	HandleError(err, "Fetching channel failed")
	jobQueue, err := channel.QueueDeclare(
		name,  // Name of the queue
		false, // Message is persisted or not
		false, // Delete message when unused
		false, // Exclusive
		false, // No Waiting time
		nil,   // Extra args
	)
	HandleError(err, "Job queue creation failed")
	return JobServer{Conn: conn, Channel: channel, Queue: jobQueue}
}

func main() {
	jobServer := getServer(queueName)
	jobServer.redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Start Workers in a go routines otherwise it would block the main function execution
	go func(conn *amqp.Connection) {
		workerProcess := Workers{
			conn: jobServer.Conn,
		}
		workerProcess.run()
	}(jobServer.Conn)

	router := mux.NewRouter()
	// Attach handlers
	router.HandleFunc("/job/database", jobServer.asyncDBHanlder)
	router.HandleFunc("/job/main", jobServer.asyncMailHandler)
	router.HandleFunc("/job/callback", jobServer.asyncCallbackHandler)
	router.HandleFunc("/job/status", jobServer.statusHandler)

	httpServer := &http.Server{
		Addr:         hostString,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(httpServer.ListenAndServe())

	// Cleanup resources
	defer jobServer.Channel.Close()
	defer jobServer.Conn.Close()
}
