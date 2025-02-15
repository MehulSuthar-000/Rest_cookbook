package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MehulSuthar-000/MetroRailAPI/dbutils"
	"github.com/emicklei/go-restful"
	_ "modernc.org/sqlite"
)

// DB Driver visible to whole program
var DB *sql.DB

// We need struct to store the data from the database

// TrainResource is a struct to map the train table in the database
type TrainResource struct {
	ID              int    `json:"id"`
	DriverName      string `json:"drivername"`
	OperatingStatus bool   `json:"operatingstatus"`
}

// StationResource is a struct to map the station table in the database
type StationResource struct {
	ID          int
	Name        string
	Openingtime time.Time
	Closingtime time.Time
}

// ScheduleResource is a struct to map the schedule table in the database
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

// Register adds paths and routes to a new service instance
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

// GET http://localhost:8000/v1/trains/1
func (t *TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("SELECT ID, DRIVER_NAME, OPERATING_STATUS FROM train WHERE ID=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		response.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
func (t *TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Printf("Request body: %+v", request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		statement, _ := DB.Prepare("INSERT INTO train (DRIVER_NAME, OPERATING_STATUS) values (?, ?)")
		result, err := statement.Exec(b.DriverName, b.OperatingStatus)
		if err == nil {
			newID, _ := result.LastInsertId()
			b.ID = int(newID)
			response.WriteHeaderAndEntity(http.StatusCreated, b)
		} else {
			response.AddHeader("Content-Type", "text/plain")
			response.WriteErrorString(http.StatusInternalServerError, err.Error())
		}
	}
}

// DELETE http://localhost:8000/v1/trains/1
func (t *TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare("DELETE FROM train WHERE ID=?")
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func main() {
	var err error
	// Connect to the database
	DB, err = sql.Open("sqlite", ".\\database\\railapi.db")
	if err != nil {
		log.Fatalf("Driver creation failed!")
	}

	// Ensure the database connection is closed when the main function ends
	defer DB.Close()

	// Create tables
	dbutils.Initialize(DB)

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Println("start listening on localhost:8000")
	server := &http.Server{
		Addr:    ":8000",
		Handler: wsContainer,
	}
	log.Fatal(server.ListenAndServe())
}
