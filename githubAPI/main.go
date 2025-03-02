package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
)

var GITHUB_TOKEN string
var requestOptions = &grequests.RequestOptions{
	Auth: []string{GITHUB_TOKEN, "x-oauth-basic"},
}

type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

func main() {
	load_env()
	var repos []Repo
	var repoUrl = "https://api.github.com/users/MehulSuthar-000/repos"
	resp := getStats(repoUrl)
	resp.JSON(&repos)
	log.Println(repos)
}

func load_env() {
	err := godotenv.Load("D:\\learn_go\\Rest_cookbook\\.env")
	if err != nil {
		log.Fatal("Unable to load env file")
	}
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
}

func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	// we are modifying the request by passing an optional RequestOptions
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	return resp
}
