package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/levigross/grequests"
	"github.com/urfave/cli"
)

var GITHUB_TOKEN string
var requestOptions *grequests.RequestOptions

func main() {
	godotenv.Load(".env")
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
	requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

	app := cli.NewApp()
	// Define commands for cli
	app.Commands = []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "Fetch the repo detail with user.\n [Usage]: githubAPI fetch user",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					// Github API logic
					var repos []Repo
					user := c.Args()[0]
					var repoUrl = fmt.Sprintf("https://api.github.com/users/%s/repos", user)
					resp := getStats(repoUrl)
					resp.JSON(&repos)
					log.Println(repos)
				} else {
					log.Println("Please give a username. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Creates a gist from the given text.\n [Usage]: githubApi name 'description' sample.txt",
			Action: func(c *cli.Context) error {
				if c.NArg() > 1 {
					// Github APi Logic
					args := c.Args()
					var postUrl = "https://api.github.com/gists"
					resp := createGist(postUrl, args)
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments. See -h to see help")
				}
				return nil
			},
		},
	}
	app.Version = "1.0.0"
	app.Run(os.Args)
}

// Struct for holding response of repositories fetch API
type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

// Struct for modelling JSON body in create Gist
type File struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

// Fetches the repos for the given Github users
func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	// you can modify the request by passing an optional
	// RequestOptions struct
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func createGist(url string, args []string) *grequests.Response {
	description := args[0]
	// remaining arguments are files names with path
	var filecontents = make(map[string]File)
	for i := 1; i < len(args); i++ {
		data, err := os.ReadFile(args[i])
		if err != nil {
			log.Println("Please check the filenames. Absolute Path (or) same directory is allowed.")
			return nil
		}
		var file File
		file.Content = string(data)
		filecontents[args[i]] = file
	}
	var gist = Gist{Description: description, Public: true, Files: filecontents}
	var postBody, _ = json.Marshal(gist)
	var requestOptions_copy = requestOptions
	requestOptions_copy.JSON = string(postBody)
	// make a Post request to GIthub
	resp, err := grequests.Post(url, requestOptions_copy)
	if err != nil {
		log.Println("Create request failed for Github API")
	}
	return resp
}
