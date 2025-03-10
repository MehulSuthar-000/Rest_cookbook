package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

type Response struct {
	License struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"license"`
}

func main() {

	err := godotenv.Load("D:\\learn_go\\Rest_cookbook\\.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	GITHUB_TOKEN := os.Getenv("GITHUB_TOKEN")

	log.Println(GITHUB_TOKEN)
	// Create a client (safe to share across requests)
	client := graphql.NewClient("https://api.github.com/graphql")

	// make a request to Github API
	req := graphql.NewRequest(`
query {
	license(key: "apache-2.0") {
		name
		description
	}
}
`)

	req.Header.Add("Authorization", "bearer "+GITHUB_TOKEN)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// run it and capture the response
	var respData Response
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	log.Println("RESPONSE: " + respData.License.Description)
}
