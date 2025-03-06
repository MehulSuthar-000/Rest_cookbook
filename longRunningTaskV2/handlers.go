package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/mehulsuthar-000/longrunningtask-v1/models"
	"github.com/streadway/amqp"
)

type JobServer struct {
	Queue       amqp.Queue
	Channel     *amqp.Channel
	Conn        *amqp.Connection
	redisClient *redis.Client
}

func (s *JobServer) publish(jsonBody []byte) error {
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonBody,
	}
	err := s.Channel.Publish(
		"",        // exchange
		queueName, // routing key(Queue)
		false,     // mandatory
		false,     // immediate
		message,
	)
	HandleError(err, "Error while generating JobID")
	return err
}

func (s *JobServer) asyncDBHanlder(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside async DB handler")
	jobId, _ := uuid.NewRandom()
	queryParams := r.URL.Query()

	// Ex: client_time: 1569174071
	unixTime, err := strconv.ParseInt(queryParams.Get("client_time"), 10, 64)
	clientTime := time.Unix(unixTime, 0)
	HandleError(err, "Error while converting client time")

	jsonBody, err := json.Marshal(models.Job{
		ID:        jobId,
		Type:      "A",
		ExtraData: models.Log{ClientTime: clientTime},
	})

	HandleError(err, "JSON body creation failed")
	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) asyncMailHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside async mail Handler")

	JobID, err := uuid.NewRandom()
	// queryParams := r.URL.Query()

	// Ex: client_time: 1569174071
	// unixTime, err := strconv.ParseInt(queryParams.Get("client_time"), 10, 64)
	// clientTime := time.Unix(unixTime, 0)
	HandleError(err, "Error while converting client time")

	jsonBody, err := json.Marshal(models.Job{ID: JobID,
		Type:      "C",
		ExtraData: "", // Can be custom data, Ex: {"email_address":
		// "packt@example.org"}
	})

	HandleError(err, "JSON body creation failed")
	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) asyncCallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside async Call back handler")
	jobID, err := uuid.NewRandom()
	// queryParams := r.URL.Query()

	// Ex: client_time: 1569174071
	// unixTime, err := strconv.ParseInt(queryParams.Get("client_time"), 10, 64)
	// clientTime := time.Unix(unixTime, 0)
	HandleError(err, "Error while converting client time")

	jsonBody, err := json.Marshal(models.Job{ID: jobID,
		Type:      "B",
		ExtraData: "", // Can be custom data, Ex: {"client_time":
		// "2020-01-22T20:38:15+02:00"}
	})

	HandleError(err, "JSON body creation failed")
	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) statusHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	uuid := queryParams.Get("uuid")
	w.Header().Set("Content-Type", "application/json")
	jobStatus := s.redisClient.Get(uuid)
	status := map[string]string{"uuid": uuid, "status": jobStatus.Val()}

	response, err := json.Marshal(status)
	HandleError(err, "Cannot Create response for client")
	w.Write(response)
}
