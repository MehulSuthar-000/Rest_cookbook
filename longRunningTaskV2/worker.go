package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mehulsuthar-000/longrunningtask-v1/models"
	"github.com/streadway/amqp"
)

type Workers struct {
	conn        *amqp.Connection
	redisClient *redis.Client
}

func (w *Workers) dbWork(job models.Job) {
	result := job.ExtraData.(map[string]interface{})
	w.redisClient.Set(job.ID.String(), "STARTED", 0)
	log.Printf("Worker %s: extracting data..., JOB: %s",
		job.Type, result)
	w.redisClient.Set(job.ID.String(), "IN PROGRESS", 0)
	time.Sleep(15 * time.Second)
	log.Printf("Worker %s: saving data to database..., JOB: %s",
		job.Type, job.ID)
	w.redisClient.Set(job.ID.String(), "DONE", 0)
}

func (w *Workers) callbackWork(job models.Job) {
	log.Printf("Worker %s: performing some long running process..., JOB: %s", job.Type, job.ID)
	time.Sleep(10 * time.Second)
	log.Printf("Worker %s: posting the data back to the given callback..., JOB: %s", job.Type, job.ID)
}
func (w *Workers) emailWork(job models.Job) {
	log.Printf("Worker %s: sending the email..., JOB: %s", job.Type, job.ID)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: sent the email successfully, JOB: %s", job.Type, job.ID)
}

func (w *Workers) run() {
	log.Println("Workers are booted up and running")
	channel, err := w.conn.Channel()
	HandleError(err, "Fetchinf channel failed")
	defer channel.Close()

	w.redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Declare a job Queue
	jobQueue, err := channel.QueueDeclare(
		queueName, // Name of the queue
		false,     // Message is persisted or not
		false,     // Delete message when unused
		false,     // Exclusive
		false,     // No Waiting time
		nil,       // Extra args
	)
	HandleError(err, "Job queue fetch failed")

	messages, _ := channel.Consume(
		jobQueue.Name, // queue
		"",            // consumer
		true,          // auto-acknowledge
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)

	go func() {
		for message := range messages {

			job := models.Job{}
			err := json.Unmarshal(message.Body, &job)

			log.Printf("Workers received a message from the queue: %s", job)
			HandleError(err, "Unable to load queue message")

			switch job.Type {
			case "A":
				w.dbWork(job)
			case "B":
				w.callbackWork(job)
			case "C":
				w.emailWork(job)
			}
		}
	}()
	defer w.conn.Close()
	wait := make(chan bool)
	<-wait // Run long-running worker
}
