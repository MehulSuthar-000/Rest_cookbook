package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int    `json:"id"`
	DriverName      string `json:"driver_name"`
	OperatingStatus bool   `json:"operating_status"`
}

// GetTrain handles GET on /trains/:id
func (c App) GetTrain() revel.Result {
	var train TrainResource
	// Getting the values from path parameters
	id := c.Params.Route.Get("train-id")
	// Setting the values for the train
	train.ID, _ = strconv.Atoi(id)
	train.DriverName = "John Doe"
	train.OperatingStatus = true
	// Returning the train as JSON
	c.Response.Status = http.StatusOK
	return c.RenderJSON(train)
}

// CreateTrain handles POST on train resource
func (c App) CreateTrain() revel.Result {
	var train TrainResource
	c.Params.BindJSON(&train)
	// Use train.DriverName and train.OperatingStatus
	// to insert into train table....
	train.ID = 2
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(train)
}

// RemoveTrain implements DELETE on train resource
func (c App) RemoveTrain() revel.Result {
	id := c.Params.Route.Get("train-id")
	// Use ID to delete record from train table....
	log.Println("Successfully deleted the resource:", id)
	c.Response.Status = http.StatusOK
	return c.RenderText("")
}

func (c App) Index() revel.Result {
	return c.Render()
}
