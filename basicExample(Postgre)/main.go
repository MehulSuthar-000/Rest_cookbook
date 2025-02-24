package main

import (
	"log"

	"github.com/mehulsuthar-000/postgreSQL/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database tables are successfully initialized")
}
